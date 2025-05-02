package dto

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	CreateAt int64  `json:"create_at"`
}

type RegisterRequest struct {
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}

type TokenResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}
