package models

type User struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	email string `json:"email"`
}
