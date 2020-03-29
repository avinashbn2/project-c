package models

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type ResourceItem struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"title"`
	URL       string    `json:"url" db:"url"`
	Tag       string    `json:"tag" db:"tag"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	// Author     string    `json:"author"`
	// Favourites uint32    `json:"favourites"`
}
type Resources []ResourceItem

func (rs Resources) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		b, err := json.Marshal(rs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return

	}
	w.WriteHeader(http.StatusBadRequest)
	http.Error(w, "Not supported", http.StatusBadRequest)

}

func (ri *ResourceItem) Save(db *sqlx.DB) error {
	query := `INSERT INTO resource_item(id, title, url, tag, created_at, updated_at) 
	VALUES(:id, :title, :url, :tag, :created_at, :updated_at)`
	_, err := db.NamedExec(query, ri)
	if err != nil {
		return err
	}
	return nil
}

func (rs *Resources) Retrieve(db *sqlx.DB) error {
	query := `Select * from resource_item`
	err := db.Select(rs, query)
	if err != nil {
		return err
	}
	return nil
}

func (r *ResourceItem) Get(w http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal("unable to convert json")
	}
	w.Write(data)
}
