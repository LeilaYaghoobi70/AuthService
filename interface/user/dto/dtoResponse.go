package dto

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
type UserResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	CreateAt int64  `json:"create_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type TokenValidationResponse struct {
	IsValid bool `isValid:"token"`
}

type SuccessfulResponse struct {
	IsValid bool `isValid:"token"`
}
