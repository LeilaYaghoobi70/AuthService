package application

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func TokenIsValid(tokenString string) (bool, error) {

	token, err := getToken(tokenString)

	if err != nil || !token.Valid {
		return false, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("invalid claims")
	}
	exp, ok := claims["exp"].(int64)
	if time.Now().After(time.Unix(exp, 0)) {
		return false, fmt.Errorf("token has expired")
	}

	return true, nil
}

func getToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return token, err
}
