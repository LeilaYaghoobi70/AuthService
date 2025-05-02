package mapper

import (
	"authService/domain/entity"
	"authService/interface/user/dto"
)

func ToUserDTO(u *entity.User) dto.User {
	return dto.User{
		Password: u.Password,
		Email:    u.Email,
		Role:     u.Role,
		CreateAt: u.CreateAt,
	}
}
