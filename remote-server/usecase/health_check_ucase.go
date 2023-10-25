package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RianNegreiros/vigilate/domain"
	"github.com/RianNegreiros/vigilate/email"
	"github.com/go-co-op/gocron"
)

type healthCheckUsecase struct {
	remoteServerRepo domain.RemoteServerRepository
	userRepo         domain.UserRepository
	contextTimeout   time.Duration
	kafkaProducer    domain.KafkaProducer
	serverStatus     map[string]bool
}

func NewHealthCheckUsecase(rsr domain.RemoteServerRepository, ur domain.UserRepository, timeout time.Duration, kafkaProducer domain.KafkaProducer) domain.HealthCheckUsecase {
	return &healthCheckUsecase{
		remoteServerRepo: rsr,
		userRepo:         ur,
		contextTimeout:   timeout,
		kafkaProducer:    kafkaProducer,
		serverStatus:     make(map[string]bool),
	}
}

func (hc *healthCheckUsecase) StartHealthChecksScheduler() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(5).Minutes().Do(hc.performServerHealthChecks)
	scheduler.StartAsync()
}

func (hc *healthCheckUsecase) performServerHealthChecks() {
	ctx := context.Background()

	servers, err := hc.remoteServerRepo.GetAll(ctx)
	if err != nil {
		log.Printf("Error getting servers: %v", err)
		return
	}

	for _, server := range servers {
		if !server.IsActive && server.NextCheckTime.After(time.Now()) {
			continue
		}

		if time.Since(server.LastCheckTime) < time.Second*5 {
			continue
		}

		go hc.checkServerStatus(ctx, server)
	}
}

func (hc *healthCheckUsecase) checkServerStatus(ctx context.Context, server domain.RemoteServer) {
	prevState, exists := hc.serverStatus[server.Address]
	err := hc.updateServerStatus(ctx, server)
	if err != nil {
		log.Printf("Error updating server status: %v", err)
		return
	}

	if !exists || prevState != server.IsActive {
		hc.serverStatus[server.Address] = server.IsActive
		if !server.IsActive {
			hc.sendNotifications(server)
		}
	}
}

func (hc *healthCheckUsecase) updateServerStatus(ctx context.Context, server domain.RemoteServer) error {
	server.IsActive = isServerUp(server.Address)
	server.LastCheckTime = time.Now()
	server.NextCheckTime = time.Now().Add(time.Minute * 5)
	err := hc.remoteServerRepo.Update(ctx, &server)
	if err != nil {
		return err
	}
	return nil
}

func (hc *healthCheckUsecase) sendNotifications(server domain.RemoteServer) {
	topic := os.Getenv("KAFKA_TOPIC")
	ctx, cancel := context.WithTimeout(context.Background(), hc.contextTimeout)
	defer cancel()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("Alert: Server Down\nServer: %s\nAddress: %s\nTimestamp: %s", server.Name, server.Address, timestamp)
	hc.kafkaProducer.SendHealthCheckResultToKafka(message, topic)

	user, err := hc.userRepo.GetUserByID(ctx, server.UserID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return
	}

	if user.NotificationPreferences.EmailEnabled {
		err = email.ResendEmailSender(user.Email, server.Name, server.Address, timestamp)

		if err != nil {
			log.Printf("Error sending email: %v", err)
			return
		}

		server.LastNotificationTime = time.Now()
		err = hc.remoteServerRepo.Update(ctx, &server)
		if err != nil {
			log.Printf("Error updating server: %v", err)
			return
		}
	}
}
