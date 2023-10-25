package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/RianNegreiros/vigilate/config"
	"github.com/RianNegreiros/vigilate/internal/database"
	"github.com/RianNegreiros/vigilate/internal/kafka"
	_remoteServerHandler "github.com/RianNegreiros/vigilate/internal/remote-server/delivery/http"
	_remoteServerRepo "github.com/RianNegreiros/vigilate/internal/remote-server/repository/postgres"
	_remoteServerUsecase "github.com/RianNegreiros/vigilate/internal/remote-server/usecase"
	_userHandler "github.com/RianNegreiros/vigilate/internal/user/delivery/http"
	_userRepo "github.com/RianNegreiros/vigilate/internal/user/repository/postgres"
	_userUsecase "github.com/RianNegreiros/vigilate/internal/user/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbConn, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Error pinging database", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("Error closing database connection", err)
		}
	}()

	pusherClient := config.NewPusherClient()

	kafkaWriterConfig := config.NewKafkaWriterConfig()
	kafkaProducer := kafka.NewKafkaProducer(kafkaWriterConfig.Brokers, kafkaWriterConfig.Topic, kafkaWriterConfig.Dialer)

	kafkaReaderConfig := config.NewKafkaReaderConfig()
	kafkaConsumer := kafka.NewKafkaConsumer(kafkaReaderConfig.Brokers, kafkaReaderConfig.Topic, kafkaReaderConfig.GroupID, kafkaWriterConfig.Dialer)

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for {
			err := kafkaConsumer.ConsumeMessages(ctx, func(message []byte) error {
				log.Println("Message received: ", string(message))
				return nil
			})
			if err != nil {
				log.Println("Error consuming messages: ", err)
			}
		}
	}()

	e := echo.New()

	contextTimeout := time.Duration(10) * time.Second

	ur := _userRepo.NewPostgresUserRepo(dbConn.GetDB())
	uu := _userUsecase.NewUserUsecase(ur, contextTimeout)
	_userHandler.NewUserHandler(e, uu)

	rsr := _remoteServerRepo.NewPostgresRemoteServerRepo(dbConn.GetDB())
	rsu := _remoteServerUsecase.NewRemoteServerUsecase(rsr, contextTimeout)
	rtm := _remoteServerUsecase.NewRealTimeMonitoringUsecase(pusherClient, rsr, contextTimeout)
	_remoteServerHandler.NewRemoteServerHandler(e, rsu, rtm)

	hcu := _remoteServerUsecase.NewHealthCheckUsecase(rsr, ur, contextTimeout, kafkaProducer)
	hcu.StartHealthChecksScheduler()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := e.Start(":8080"); err != nil {
			log.Println("Error starting server: ", err)
		}
	}()

	wg.Wait()
}
