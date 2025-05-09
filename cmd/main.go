package main

import (
	"authService/application/auth"
	"authService/application/user"
	"authService/infrastructure/db"
	"authService/infrastructure/repository"
	user2 "authService/interface/user"
	"log"
)

func main() {
	db.Connect()

	if err := db.Close(); err != nil {
		log.Fatalf("error closing db: %v", err)
	}

	if err := db.CreateSchema(db.DB); err != nil {
		log.Fatalf("error closing db: %v", err)
	}

	userRepo := repository.UserRepository(db.DB)
	userAuth := auth.AuthService()
	userService := user.UserService(userRepo, userAuth)
	handler := user2.RouterHandler(userService)
	user2.RegisterRoutes(handler)
}
