package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection(conStr string) *sqlx.DB {

	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
