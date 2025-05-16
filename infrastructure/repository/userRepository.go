package repository

import (
	"authService/domain/entity"
	"authService/domain/user"
	"authService/infrastructure/db"
	"authService/mapper"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2/log"
)

type userRepository struct {
	db *pg.DB
}

func UserRepository(db *pg.DB) user.Repository {
	return userRepository{db: db}
}

func (r userRepository) FindUserByEmail(email string) (*entity.User, error) {
	u := &db.User{}

	err := r.db.Model(u).Where("email = ?", email).Select()
	if err != nil {
		return nil, err
	}

	return mapper.ToUserDomain(u), nil
}

func (r userRepository) Signup(email, password string) error {
	u := &db.User{
		Email:    email,
		Password: password,
	}
	resut, err := r.db.Model(u).Insert()
	log.Infof("resut : %v", resut.Model())
	return err

}
