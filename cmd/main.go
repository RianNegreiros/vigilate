package main

import (
	"log"
	"time"

	"github.com/RianNegreiros/vigilate/config"
	"github.com/RianNegreiros/vigilate/infra/database"
	_remoteServerHandler "github.com/RianNegreiros/vigilate/remote-server/delivery/http"
	_remoteServerRepo "github.com/RianNegreiros/vigilate/remote-server/repository/postgres"
	_remoteServerUsecase "github.com/RianNegreiros/vigilate/remote-server/usecase"
	_userHandler "github.com/RianNegreiros/vigilate/user/delivery/http"
	_userRepo "github.com/RianNegreiros/vigilate/user/repository/postgres"
	_userUsecase "github.com/RianNegreiros/vigilate/user/usecase"
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

	config.NewPusherClient()

	e := echo.New()

	contextTimeout := time.Duration(5) * time.Second

	ur := _userRepo.NewPostgresUserRepo(dbConn.GetDB())
	uu := _userUsecase.NewUserUsecase(ur, contextTimeout)
	_userHandler.NewUserHandler(e, uu)

	rsr := _remoteServerRepo.NewPostgresRemoteServerRepo(dbConn.GetDB())
	rsu := _remoteServerUsecase.NewRemoteServerUsecase(rsr, contextTimeout)
	rsu.StartServerHealthCheck()
	_remoteServerHandler.NewRemoteServerHandler(e, rsu)

	log.Fatal(e.Start(":8080"))
}
