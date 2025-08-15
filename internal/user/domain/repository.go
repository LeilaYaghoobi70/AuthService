package domain

import (
	"authService/internal/user/domain/entity"
	"context"
)

type Repository interface {
	FindUserByEmail(email string) (*entity.User, error)
	FindUserByUsername(username string) (*entity.User, error)
	Signup(username, email, password string) error
	Authenticate(username, email, password string, ctx context.Context) (bool, error)
}
