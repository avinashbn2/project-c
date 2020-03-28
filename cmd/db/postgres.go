package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection() *sqlx.DB {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
