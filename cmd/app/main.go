package main

import (
	"cproject/cmd/db"
	"cproject/internal/models"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type DB *sqlx.DB

type App struct {
	db DB
}

var app App
var allResources models.Resources

func main() {

	// TODO pass connection variableshere
	db := db.NewConnection()
	app = App{db}
	err := models.AddResourceItem(app.db)
	if err != nil {
		log.Fatal(err)
	}
	// allResources = resources
	// http.Handle("/", allResources)
	err = http.ListenAndServe(":3001", nil)
	fmt.Println(err)
}
