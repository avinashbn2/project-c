// This is going to be the main repository where new models can be associated with stuff!
package repository

import (
	"cproject/internal/models"
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type TagRepo struct {
	db *sqlx.DB
}

func NewTagRepository(db *sqlx.DB) *TagRepo {
	return &TagRepo{
		db,
	}
}

func (rp *TagRepo) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tag := &models.Tag{}
		err := json.NewDecoder(r.Body).Decode(tag)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = tag.Save(rp.db)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		rsp, err := json.Marshal(tag)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte(rsp))

	}
}
