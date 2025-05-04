package user

import (
	"authService/application/auth"
	"authService/domain/entity"
	"authService/domain/user"
	"errors"
	"github.com/go-pg/pg/v10"
)

type Service interface {
	FindUserByEmail(email string) (*entity.User, error)
	Login(email, password string) (string, error)
	Signup(email, password string) (string, error)
}

type userService struct {
	repo user.Repository
	auth auth.Service
}

func UserService(repo user.Repository, auth auth.Service) Service {
	return &userService{
		repo: repo,
		auth: auth,
	}
}

func (u userService) FindUserByEmail(email string) (*entity.User, error) {
	user, err := u.repo.FindUserByEmail(email)
	return user, err
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

func (u userService) Signup(email, password string) (string, error) {
	user, err := u.repo.FindUserByEmail(email)
	if err != nil && !errors.Is(err, pg.ErrNoRows) {
		return "", errors.New("something went wrong")
	}
	if user != nil {
		return "", errors.New("user already exists")
	}

	token, err := u.auth.GenerateToken(email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	err = u.repo.Signup(email, password)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
