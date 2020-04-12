package main

import (
	"cproject/cmd/db"
	"cproject/cmd/repository"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type App struct {
	db *sqlx.DB
}

// var app App
// var allResources models.Resources

// func (app *App) InitRouter() *chi.Mux {
// 	r := chi.NewRouter()
// 	// res := &models.ResourceItem{
// 	// 	ID:        "9237051",
// 	// 	Name:      "Go",
// 	// 	URL:       "http://Go2.com",
// 	// 	Tag:       "Go",
// 	// 	CreatedAt: time.Now().UTC(),
// 	// 	UpdatedAt: time.Now().UTC(),
// 	// }
// 	r.Route("/resources", func(router chi.Router) {
// 		router.Get("/", res.Get(app.db))
// 		router.Get("/res", res.Get(app.db))

// 	})
// 	return r

// }

func main() {

	// TODO pass connection variableshere
	database := db.NewConnection()
	// resources, err := models.GetResourceItem(app.db)
	// fmt.Println(resources)
	// var resources models.Resources
	// err := resources.Retrieve(db)
	// fmt.Println(resources)
	// allResources = resource5ks
	// http.Handle("/", allResources)
	// mux := app.InitRouter()
	repo := repository.NewResourceRepo(database)

	server := NewServer(":3000", repo)
	err := http.ListenAndServe(":3001", server.Handler())
	if err != nil {
		fmt.Print(err)
	}
	// fmt.Println(err)
}

// ritem := &models.ResourceItem{
// 	ID:        "9237055",
// 	Name:      "React",
// 	URL:       "http://React2.com",
// 	Tag:       "React",
// 	CreatedAt: time.Now().UTC(),
// 	UpdatedAt: time.Now().UTC(),
// }
// err := ritem.Save(database)
// if err != nil {
// 	fmt.Print(err)
// }
