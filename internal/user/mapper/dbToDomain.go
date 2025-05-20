package mapper

import (
	"authService/internal/user/domain/entity"
	"authService/internal/user/infrastructure/db"
)

func ToUserDomain(u *db.User) *entity.User {
	return &entity.User{
		Password: u.Password,
		Email:    u.Email,
		Role:     u.Role,
		CreateAt: u.CreateAt,
	}
}
