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

// PostTags godoc
type PostTags struct {
	Tags []int `json:"tags"`
}

// PostLikes godoc
type PostLikes struct {
	Likes int `json:"likes"`
}

// PostLike godoc
type PostLike struct {
	UserID int  `json:"user_id"`
	Like   bool `json:"like"`
}

//PostComment godoc
type PostComment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Guest     bool      `json:"guest"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UserID    int       `json:"user_id"`
	Comment   string    `json:"comment"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// PostModel godoc
type PostModel struct{}
