package main

import (
	"authService/application/auth"
	"authService/application/user"
	"authService/infrastructure/db"
	"authService/infrastructure/repository"
	user2 "authService/interface/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
		return
	}
	db.Connect()
	defer db.Close()
	if err := db.CreateSchema(db.DB); err != nil {
		log.Fatalf("error closing db: %v", err)
	}

	userRepo := repository.UserRepository(db.DB)
	userAuth := auth.AuthService()
	userService := user.UserService(userRepo, userAuth)
	handler := user2.RouterHandler(userService)
	app := fiber.New()
	app.Use(logger.New())
	user2.RegisterRoutes(app, handler)
	app.Listen("127.0.0.1:" + os.Getenv("GOLANG_PORT"))
}
