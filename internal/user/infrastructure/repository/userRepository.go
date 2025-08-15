package repository

import (
	"authService/internal/user/domain"
	"authService/internal/user/domain/entity"
	"authService/internal/user/infrastructure/db"
	"authService/internal/user/mapper"
	"context"
	"database/sql"
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *pg.DB
}

func UserRepository(db *pg.DB) domain.Repository {
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

func (r userRepository) Authenticate(username, email, password string, ctx context.Context) (bool, error) {
	const q = `SELECT password_hash FROM users WHERE username = $1 AND email = $1 LIMIT 1`
	u := &db.User{
		Username: username, Email: email,
	}
	var hash string
	_, err := r.db.QueryOneContext(ctx, q, u)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return false, nil
	}
	return true, nil
}

func (r userRepository) FindUserByUsername(username string) (*entity.User, error) {
	u := &db.User{}

	err := r.db.Model(u).Where("username = ?", username).Select()
	if err != nil {
		return nil, err
	}

	return mapper.ToUserDomain(u), nil
}

func (r userRepository) Signup(username, email, password string) error {
	u := &db.User{
		Email:    email,
		Password: password,
		Username: username,
	}
	resut, err := r.db.Model(u).Insert()
	log.Infof("resut : %v", resut.Model())
	return err

}
