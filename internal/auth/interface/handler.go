package _interface

import (
	"authService/internal/auth/application"
	"authService/internal/user/interface/dto"

	"authService/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	auth application.Service
}

func RouterHandler(s application.Service) Handler {
	return Handler{
		auth: s,
	}
}

func (h *Handler) ValidationToken(c *fiber.Ctx) error {
	isValid, err := h.auth.IsValidateToken(c.GetRespHeader("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(dto.Response{Status: fiber.StatusOK, Data: dto.TokenValidationResponse{IsValid: isValid}})
}
