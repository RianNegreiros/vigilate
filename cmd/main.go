package main

import (
	"log"
	"time"

	"github.com/RianNegreiros/vigilate/config"
	"github.com/RianNegreiros/vigilate/framework/kafka"
	"github.com/RianNegreiros/vigilate/infra/database"
	_remoteServerHandler "github.com/RianNegreiros/vigilate/remote-server/delivery/http"
	_remoteServerRepo "github.com/RianNegreiros/vigilate/remote-server/repository/postgres"
	_remoteServerUsecase "github.com/RianNegreiros/vigilate/remote-server/usecase"
	_userHandler "github.com/RianNegreiros/vigilate/user/delivery/http"
	_userRepo "github.com/RianNegreiros/vigilate/user/repository/postgres"
	_userUsecase "github.com/RianNegreiros/vigilate/user/usecase"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
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
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	kafkaConfig := config.NewKafkaConfig()
	producer, err := ckafka.NewProducer(kafkaConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	kafkaProducer := kafka.NewKafkaProducer(producer)

	consumer, err := ckafka.NewConsumer(kafkaConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	config.NewPusherClient()

	e := echo.New()

	contextTimeout := time.Duration(5) * time.Second

	ur := _userRepo.NewPostgresUserRepo(dbConn.GetDB())
	uu := _userUsecase.NewUserUsecase(ur, contextTimeout)
	_userHandler.NewUserHandler(e, uu)

	rsr := _remoteServerRepo.NewPostgresRemoteServerRepo(dbConn.GetDB())
	rsu := _remoteServerUsecase.NewRemoteServerUsecase(rsr, contextTimeout)
	_remoteServerHandler.NewRemoteServerHandler(e, rsu)

	hcu := _remoteServerUsecase.NewHealthCheckUsecase(rsr, contextTimeout, kafkaProducer)
	hcu.StartHealthChecksScheduler()

	log.Fatal(e.Start(":8080"))
}
