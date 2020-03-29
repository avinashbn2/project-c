package models

type Tag struct {
	TID    string `json:"tid" db:"tid"`
	Name   string `json:"name" db:"name"`
	Parent string `json:"parent" db:"parent"`
}
