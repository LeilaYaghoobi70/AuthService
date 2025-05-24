package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	DbPort     string
	DbPassword string
	DbHost     string
	DbName     string
	GolangPort string
)

func LoadEvn() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	DbPort = os.Getenv("DB_PORT")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost = os.Getenv("DB_HOST")
	DbName = os.Getenv("DB_NAME")
	GolangPort = os.Getenv("GOLANG_PORT")
	return nil
}
