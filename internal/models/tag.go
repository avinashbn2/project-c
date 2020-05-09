package models

import "github.com/jmoiron/sqlx"

type Tag struct {
	TID    string `json:"tid" db:"tid"`
	Name   string `json:"name" db:"name"`
	Parent string `json:"parent" db:"parent"`
}

func (ri *Tag) Save(db *sqlx.DB) error {
	query := `INSERT INTO tag(name) VALUES(:name)`
	_, err := db.NamedExec(query, ri)
	if err != nil {
		return err
	}
	return nil
}
