package main

import (
	"log"

	"github.com/RianNegreiros/vigilate/db"
	"github.com/RianNegreiros/vigilate/internal/service"
	"github.com/RianNegreiros/vigilate/internal/user"
	"github.com/labstack/echo"
)

func main() {
	dbConn, err := db.NewDatabase()
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

	userRepo := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepo)
	user.NewUserHandler(e, userService)

	serviceRepo := service.NewRepository(dbConn.GetDB())
	serviceService := service.NewService(serviceRepo)
	service.NewServiceHandler(e, serviceService)

	log.Fatal(e.Start(":8080"))
}
