package repository

import (
	"cproject/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
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
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		bytes, err := json.Marshal(&rs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(bytes))
	}
}

//Add : Post Resource data to table
func (rp *ResourceRepo) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resource := &models.ResourceItem{}
		err := json.NewDecoder(r.Body).Decode(resource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = resource.Save(rp.db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		rsp, err := json.Marshal(resource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(rsp))
	}
}

//FindByID : Select  from the table by id
func (rp *ResourceRepo) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		fmt.Printf("%T %s", id, id)

		query := fmt.Sprintf("Select * from resource_item where id=%s", string(id))
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
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(bytes))
	}
}

//Update : Update Resource data
func (rp *ResourceRepo) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		query := fmt.Sprintf("Select * from resource_item where id=%d limit 1", id)
		var resources models.Resources
		err = rp.db.Select(&resources, query)
		if err != nil {
			log.Fatal(err, err.Error(), "update")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var rs models.ResourceItem
		rs.Override(resources[0])
		err = json.NewDecoder(r.Body).Decode(&rs)

		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rs.UpdatedAt = time.Now().UTC()
		err = rs.Update(rp.db)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		rsp, err := json.Marshal(&rs)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(rsp))
	}
}
