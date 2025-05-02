package user

import (
	"authService/application/user"
	"authService/interface/user/dto"
	"authService/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service user.Service
}

var validate = validator.New()

func RouterHandler(s user.Service) Handler {
	return Handler{service: s}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var login dto.LoginRequest

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.BadRequest("invalid request"))
	}

	if err := validate.Struct(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}

	token, err := h.service.Login(login.Email, login.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalError("service error"))
	}

	return c.Status(fiber.StatusOK).JSON(dto.TokenResponse{Status: fiber.StatusOK, Token: token})
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) FindUserByEmail(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
