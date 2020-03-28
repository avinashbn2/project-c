package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type ResourceItem struct {
	ID        string    `json:"id"`
	Name      string    `json:"title"`
	URL       string    `json:"url"`
	Tag       Tag       `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func AddResourceItem(db *sqlx.DB) error {
	tx := db.MustBegin()
	_, err := tx.Exec("INSERT INTO resource_item (id, title, url, tag, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", "109293023", "gogithub", "http://go.githug.com", "92809324", time.Now().UTC(), time.Now().UTC())
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}
