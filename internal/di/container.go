package di

import (
	authApplication "authService/internal/auth/application"
	"authService/internal/auth/handler"
	"authService/internal/user/application"
	"authService/internal/user/domain"
	userHanlder "authService/internal/user/handler"
	"authService/internal/user/infrastructure/db"
	"authService/internal/user/infrastructure/repository"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

type container struct {
	DB          *pg.DB
	App         *fiber.App
	Repo        domain.Repository
	AuthService authApplication.Service
	AuthHandler handler.Handler
	UserHandler userHanlder.Handler
	UserService application.Service
}

type Container *container

func NewContainer() (Container, error) {
	db.Connect()
	defer db.Close()
	if err := db.CreateSchema(db.DB); err != nil {
		log.Fatalf("error closing db: %v", err)
		return nil, err
	}

	userRepo := repository.UserRepository(db.DB)

	authService := authApplication.AuthService()
	userService := application.UserService(userRepo, authService)

	userHandler := userHanlder.RouterHandler(userService)
	authHandler := handler.RouterHandler(authService)

	app := fiber.New()
	app.Use(logger.New())

	return &container{
		DB:          db.DB,
		App:         app,
		Repo:        userRepo,
		AuthService: authService,
		AuthHandler: authHandler,
		UserHandler: userHandler,
		UserService: userService,
	}, nil
}
