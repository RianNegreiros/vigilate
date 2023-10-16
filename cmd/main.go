package main

import (
	"log"
	"time"

	"github.com/RianNegreiros/vigilate/infra/database"
	_userHandler "github.com/RianNegreiros/vigilate/user/delivery/http"
	_userRepo "github.com/RianNegreiros/vigilate/user/repository/postgres"
	_userUsecase "github.com/RianNegreiros/vigilate/user/usecase"

	"github.com/labstack/echo"
)

func main() {
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

	e := echo.New()

	contextTimeout := time.Duration(5) * time.Second

	ur := _userRepo.NewPostgresUserRepo(dbConn.GetDB())
	uu := _userUsecase.NewUserUsecase(ur, contextTimeout)
	_userHandler.NewUserHandler(e, uu)

	log.Fatal(e.Start(":8080"))
}
