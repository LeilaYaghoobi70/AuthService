package mapper

import (
	"authService/domain/entity"
	"authService/interface/user/dto"
)

func ToUserDTO(u *entity.User) dto.UserResponse {
	return dto.UserResponse{
		Password: u.Password,
		Email:    u.Email,
		Role:     u.Role,
		CreateAt: u.CreateAt,
	}
}
