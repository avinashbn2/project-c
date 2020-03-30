package main

import (
	"cproject/cmd/db"
	"cproject/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

type App struct {
	db *sqlx.DB
}

// var app App
// var allResources models.Resources

func (app *App) InitRouter() *chi.Mux {
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
		router.Get("/", res.Get(app.db))
		router.Get("/res", res.Get(app.db))

	})
	return r
}
func main() {

	// TODO pass connection variableshere
	db := db.NewConnection()
	app := &App{db}
	// resources, err := models.GetResourceItem(app.db)
	// fmt.Println(resources)
	// var resources models.Resources
	// err := resources.Retrieve(db)
	// fmt.Println(resources)
	// allResources = resources
	// http.Handle("/", allResources)
	mux := app.InitRouter()
	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(err)
}

// ritem := &models.ResourceItem{
// 	ID:        "9237053",
// 	Name:      "Rust",
// 	URL:       "http://Rust2.com",
// 	Tag:       "Rust",
// 	CreatedAt: time.Now().UTC(),
// 	UpdatedAt: time.Now().UTC(),
// }
// err := ritem.Save(db)
