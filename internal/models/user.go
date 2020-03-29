package models

type User struct {
	UID   string `json:"uid" db:"uid"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}
