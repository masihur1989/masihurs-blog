package posts

import "time"

// Post godoc
type Post struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	UserID     int       `json:"user_id"`
	CategoryID int       `json:"category_id"`
	PostView   int       `json:"post_view"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// PostView godoc
type PostView struct {
	CurrentView int `json:"post_view"`
}

// PostModel godoc
type PostModel struct{}
