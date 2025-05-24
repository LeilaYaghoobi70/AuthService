package main

import (
	authHandler "authService/internal/auth/handler"
	"authService/internal/config"
	"authService/internal/di"
	userHandler "authService/internal/user/handler"
	"log"
)

func init() {
	err := config.LoadEvn()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
		return
	}
}

func main() {
	container, err := di.NewContainer()

	if err != nil {
		log.Fatal(err)
		return
	}

	api := container.App.Group("/api/v1")

	userHandler.RegisterRoutes(api, container.UserHandler)
	authHandler.RegisterRoutes(api, container.AuthHandler)

	container.App.Listen(":" + config.GolangPort)
}
