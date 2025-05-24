package handler

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router, h Handler) {
	registerVersionOnesApi(h, app)
}

func registerVersionOnesApi(h Handler, api fiber.Router) {
	api.Get("validate", h.ValidationToken)
}
