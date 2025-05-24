package mapper

import (
	"authService/internal/user/domain/entity"
	"authService/internal/user/handler/dto"
)

func ToUserDTO(u *entity.User) dto.UserResponse {
	return dto.UserResponse{
		Password: u.Password,
		Email:    u.Email,
		Role:     u.Role,
		CreateAt: u.CreateAt,
	}
}
