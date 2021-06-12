package models

type User struct {
	ID    int    `json:"id" db:"uid"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}
