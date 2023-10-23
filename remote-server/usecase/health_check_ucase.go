package usecase

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), hc.contextTimeout)
	defer cancel()

	servers, err := hc.remoteServerRepo.GetAll(ctx)
	if err != nil {
		return
	}

	for _, server := range servers {
		go func(server domain.RemoteServer) {
			server.IsActive = isServerUp(server.Address)
			server.LastCheckTime = time.Now()
			server.NextCheckTime = time.Now().Add(time.Minute * 5)
			err = hc.remoteServerRepo.Update(ctx, &server)
			if err != nil {
				return
			}

			var result string
			if server.IsActive {
				result = "up"
			} else {
				result = "down"
			}

			topic := "health-check-results"
			hc.kafkaProducer.SendHealthCheckResultToKafka(result, topic)
		}(server)
	}
}
