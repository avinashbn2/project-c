package resource

import "time"

type ResourceItem struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"title"`
	URL       string    `json:"url" db:"url"`
	Tag       string    `json:"tag" db:"tag"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	SiteName  string    `json:"sitename" db:"sitename"`
	Excerpt   string    `json:"excerpt" db:"excerpt"`
	Author    string    `json:"author" db:"author"`
	Image     string    `json:"image" db:"image"`
	Likes     int       `json:"likes" db:"likes"`
	UserLike  bool      `json:"likeByUser" db:"like_by_user"`
}

type Resources []ResourceItem
