package auth

import (
	"authService/domain/entity"
)

type Service interface {
	GenerateToken(email string) (string, error)
	GetUserByToken(token string) (entity.User, error)
	IsValidateToken(token string) (bool, error)
}

type service struct {
}

func AuthService() Service {
	return service{}
}

func (s service) GenerateToken(email string) (string, error) {
	return GenerateToken(email)
}
func (s service) IsValidateToken(token string) (bool, error) {
	return TokenIsValid(token)
}

func (s service) GetUserByToken(token string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}
