package models

type Tag struct {
	TID    string `json:"tid"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}
