package handler

import (
	"authService/internal/user/application"
	"authService/internal/user/handler/dto"
	"authService/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userService application.Service
}

var validate = validator.New()

func RouterHandler(s application.Service) Handler {
	return Handler{
		userService: s,
	}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var login dto.LoginRequest

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.BadRequest("invalid request"))
	}

	if err := validate.Struct(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}

	token, err := h.userService.Login(login.Email, login.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalError("application error"))
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{Status: fiber.StatusOK, Data: dto.TokenResponse{Token: token}})
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	var registerRequest dto.RegisterRequest

	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.BadRequest("invalid request"))
	}

	err := h.userService.Signup(registerRequest.Email, registerRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(dto.Response{Status: fiber.StatusOK, Data: ""})
}
