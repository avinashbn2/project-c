package models

type UserFavourites struct {
	UID string `json:"uid" db:"uid"`
	RID string `json:"rid" db:"rid"`
}
