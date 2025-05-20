package _interface

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, h Handler) {
	registerVersionOnesApi(h, app)
}

func registerVersionOnesApi(h Handler, app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/login", h.Login)
	api.Post("/signUp", h.Signup)
}
