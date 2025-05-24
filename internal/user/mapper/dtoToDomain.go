package mapper

import (
	"authService/internal/user/domain/entity"
	"authService/internal/user/handler/dto"
)

func ToDomainUser(dto dto.UserResponse) *entity.User {
	return &entity.User{
		Password: dto.Password,
		Email:    dto.Email,
		Role:     dto.Role,
		CreateAt: dto.CreateAt,
	}
}
