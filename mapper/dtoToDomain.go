package mapper

import (
	"authService/domain/entity"
	"authService/interface/user/dto"
)

func ToDomainUser(dto dto.UserResponse) *entity.User {
	return &entity.User{
		Password: dto.Password,
		Email:    dto.Email,
		Role:     dto.Role,
		CreateAt: dto.CreateAt,
	}
}
