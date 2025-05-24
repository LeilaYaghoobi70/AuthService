package db

import (
	"authService/internal/config"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
)

var DB *pg.DB

func Connect() {
	opts := pg.Options{
		Addr:     config.DbHost + ":" + config.DbPort,
		Password: config.DbPassword,
		Database: config.DbName,
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
