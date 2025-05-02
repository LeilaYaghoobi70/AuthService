package errors

import "github.com/gofiber/fiber/v2"

type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func BadRequest(message string) *AppError {
	return &AppError{
		Status:  fiber.StatusBadRequest,
		Message: message,
	}
}

func NewNotFound(message string) *AppError {
	return &AppError{
		Status:  fiber.StatusNotFound,
		Message: message,
	}
}

func InternalError(message string) *AppError {
	return &AppError{
		Status:  fiber.StatusInternalServerError,
		Message: message,
	}
}
