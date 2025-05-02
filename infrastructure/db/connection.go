package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func Connect() {
	opts := pg.Options{
		Addr:     "localhost:5432",
		Database: "authService",
	}
	DB = pg.Connect(&opts)
}

func Close() error {
	if DB == nil {
		return nil
	}
	return DB.Close()
}

func CreateSchema(db *pg.DB) error {
	return db.Model((*User)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
}
