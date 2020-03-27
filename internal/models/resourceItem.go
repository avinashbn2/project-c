package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResourceItem struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Tag        Tag `json:"tag"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	Author     string `json:"author"`
	Favourites uint32 `json:"favourites"`
}
type Resources []ResourceItem

func (rs Resources) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(rs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type","application/json")
		w.Write(b)

}
