package user

import (
	"authService/domain/entity"
)

type Repository interface {
	FindUserByEmail(email string) (*entity.User, error)
	Signup(email, password string) error
}
