package main

import (
	"log"

	"github.com/RianNegreiros/vigilate/db"
	"github.com/RianNegreiros/vigilate/internal/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	e.Use(middleware.CSRF())

	userRepo := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepo)
	user.NewUserHandler(e, userService)

	log.Fatal(e.Start(":8080"))
}
