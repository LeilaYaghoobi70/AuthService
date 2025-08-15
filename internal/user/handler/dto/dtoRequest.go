package dto

type RegisterRequest struct {
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=6"`
}

type LoginRequest struct {
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}
