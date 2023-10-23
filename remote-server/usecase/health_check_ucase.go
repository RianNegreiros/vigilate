package usecase

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/RianNegreiros/vigilate/domain"
	"github.com/go-co-op/gocron"
)

type healthCheckUsecase struct {
	remoteServerRepo domain.RemoteServerRepository
	contextTimeout   time.Duration
	kafkaProducer    domain.KafkaProducer
}

func NewHealthCheckUsecase(r domain.RemoteServerRepository, timeout time.Duration, kafkaProducer domain.KafkaProducer) domain.HealthCheckUsecase {
	return &healthCheckUsecase{
		remoteServerRepo: r,
		contextTimeout:   timeout,
		kafkaProducer:    kafkaProducer,
	}
}

func (hc *healthCheckUsecase) StartHealthChecksScheduler() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(5).Seconds().Do(hc.performServerHealthChecks)
	scheduler.StartAsync()
}

func (hc *healthCheckUsecase) performServerHealthChecks() {
	ctx := context.Background()

	servers, err := hc.remoteServerRepo.GetAll(ctx)
	if err != nil {
		log.Println("Error getting all servers: ", err)
		return
	}

	for _, server := range servers {
		go hc.checkServerStatus(ctx, server)
	}
}

func (hc *healthCheckUsecase) checkServerStatus(ctx context.Context, server domain.RemoteServer) {
	server.IsActive = isServerUp(server.Address)
	server.LastCheckTime = time.Now()
	server.NextCheckTime = time.Now().Add(time.Second * 5)
	err := hc.remoteServerRepo.Update(ctx, &server)
	if err != nil {
		log.Println("Error updating server: ", err)
		return
	}

	var result string
	if server.IsActive {
		result = "up"
	} else {
		result = "down"
	}

	topic := os.Getenv("KAFKA_TOPIC")
	hc.kafkaProducer.SendHealthCheckResultToKafka(result, topic)
}
