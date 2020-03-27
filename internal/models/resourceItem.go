package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResourceItem struct {
	ID         string `json:id`
	Name       string
	URL        string
	Tag        Tag
	Created    time.Time
	Updated    time.Time
	Author     string
	Favourites uint32
}
type Resources []ResourceItem

func (rs Resources) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(rs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(b)
}
