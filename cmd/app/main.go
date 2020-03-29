package main

import (
	"cproject/cmd/db"
	"cproject/internal/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

type DB *sqlx.DB

type App struct {
	db DB
}

var app App
var allResources models.Resources

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	res := &models.ResourceItem{
		ID:        "9237051",
		Name:      "Go",
		URL:       "http://Go2.com",
		Tag:       "Go",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	r.Route("/resources", func(router chi.Router) {
		router.Get("/", res.Get)
		router.Get("/res", res.Get)

	})
	return r
}
func main() {

	// TODO pass connection variableshere
	db := db.NewConnection()
	app = App{db}
	// resources, err := models.GetResourceItem(app.db)
	// fmt.Println(resources)
	// ritem := &models.ResourceItem{
	// 	ID:        "9237051",
	// 	Name:      "Go",
	// 	URL:       "http://Go2.com",
	// 	Tag:       "Go",
	// 	CreatedAt: time.Now().UTC(),
	// 	UpdatedAt: time.Now().UTC(),
	// }
	// err := ritem.Save(db)
	var resources models.Resources
	err := resources.Retrieve(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resources)
	// allResources = resources
	// http.Handle("/", allResources)
	mux := InitRouter()
	err = http.ListenAndServe(":3001", mux)
	fmt.Println(err)
}
