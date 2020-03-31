package repository

import (
	"cproject/internal/models"
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type ResourceRepo struct {
	db *sqlx.DB
}

func NewResourceRepo(database *sqlx.DB) *ResourceRepo {
	return &ResourceRepo{
		db: database,
	}
}

//FindAll : Select all from the table
func (rp *ResourceRepo) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := `Select * from resource_item`
		var rs models.Resources
		err := rp.db.Select(&rs, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		bytes, err := json.Marshal(&rs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(bytes))
	}
}
