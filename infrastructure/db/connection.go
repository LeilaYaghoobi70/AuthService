package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"os"
)

var DB *pg.DB

func Connect() {
	opts := pg.Options{
		Addr:     ":" + os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		Database: os.Getenv("DB_NAME"),
	}

	DB = pg.Connect(&opts)
}

func Close() {
	if DB == nil {
		log.Print("database is not init ")
		return
	}
	err := DB.Close()
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}

func CreateSchema(db *pg.DB) error {
	return db.Model((*User)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
}
