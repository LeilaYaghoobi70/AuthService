package handler

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(api fiber.Router, h Handler) {
	registerVersionOnesApi(h, api)
}

func registerVersionOnesApi(h Handler, api fiber.Router) {
	api.Post("/login", h.Login)
	api.Post("/sign", h.Signup)
}
