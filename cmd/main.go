package main

import (
	authApplication "authService/internal/auth/application"
	authInterface "authService/internal/auth/interface"
	"authService/internal/user/application"
	"authService/internal/user/infrastructure/db"
	"authService/internal/user/infrastructure/repository"
	userInterface "authService/internal/user/interface"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
		return
	}
}

func main() {
	db.Connect()
	defer db.Close()
	if err := db.CreateSchema(db.DB); err != nil {
		log.Fatalf("error closing db: %v", err)
	}

	userRepo := repository.UserRepository(db.DB)

	authService := authApplication.AuthService()
	userService := application.UserService(userRepo, authService)

	userHandler := userInterface.RouterHandler(userService)
	authHandler := authInterface.RouterHandler(authService)

	app := fiber.New()
	app.Use(logger.New())

	userInterface.RegisterRoutes(app, userHandler)
	authInterface.RegisterRoutes(app, authHandler)

	app.Listen("127.0.0.1:" + os.Getenv("GOLANG_PORT"))
}
