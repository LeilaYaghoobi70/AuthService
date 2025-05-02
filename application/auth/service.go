package auth

import (
	"authService/domain/entity"
)

type Service interface {
	GenerateToken(email string, password string) (string, error)
	GetUserByToken(token string) (entity.User, error)
}

type service struct {
}

func AuthService() Service {
	return service{}
}

func (s service) GenerateToken(email string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetUserByToken(token string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}
