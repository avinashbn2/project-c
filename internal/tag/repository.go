package tag

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *repository {
	return &repository{
		db,
	}
}

func (rp *repository) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//tag := &models.Tag{}
		//err := json.NewDecoder(r.Body).Decode(tag)

		//if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		//return
		//}

		//err = tag.Save(rp.db)

		//if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		//return
		//}

		//w.Header().Add("Content-Type", "application/json")
		//rsp, err := json.Marshal(tag)

		//if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		//return
		//}
		//w.Write([]byte(rsp))
		w.Write([]byte("test"))

	}
}

func (ri *Tag) Save(db *sqlx.DB) error {
	query := `INSERT INTO tag(name) VALUES(:name)`
	_, err := db.NamedExec(query, ri)
	if err != nil {
		return err
	}
	return nil
}
