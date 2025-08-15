package application

import (
	authApplication "authService/internal/auth/application"
	"authService/internal/user/domain"
	"errors"
)

type Service interface {
	TokenIsValid(email string) (bool, error)
	Login(email, password string) (string, error)
	Signup(username, email, password string) error
}

type userService struct {
	repo domain.Repository
	auth authApplication.Service
}

func UserService(repo domain.Repository, auth authApplication.Service) Service {
	return &userService{
		repo: repo,
		auth: auth,
	}
}

func (u userService) TokenIsValid(token string) (bool, error) {
	return authApplication.TokenIsValid(token)
}

func (u userService) Login(email, password string) (string, error) {
	user, err := u.repo.FindUserByEmail(email)

	if err != nil {
		return "", errors.New("user not found")
	}

	if password != user.Password {
		return "", errors.New("password incorrect")
	}

	token, err := u.auth.GenerateToken(email)

	if err != nil {
		return "", errors.New("something went wrong")
	}

	return token, nil
}

func (u userService) Signup(username, email, password string) error {
	_, err := u.repo.Authenticate(username, email, password)
	if err != nil {
		return err
	}
	err = u.repo.Signup(username, email, password)
	if err != nil {
		return errors.New("failed to generate token")
	}

	return nil
}
