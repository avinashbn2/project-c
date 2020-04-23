package main

import (
	"cproject/cmd/db"
	"cproject/cmd/repository"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
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
var once sync.Once

func loadConfig(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal("Invlaid type assertion")
	}

	return value
}
func main() {
	connString := loadConfig("DB_CONN_STRING")
	// TODO pass connection variableshere
	database := db.NewConnection(connString)

	repo := repository.NewResourceRepo(database)

	server := NewServer(":3000", repo)
	err := http.ListenAndServe(":3001", server.Handler())
	if err != nil {
		fmt.Print(err)
	}
}
