package users

import "time"

// User data model
type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	RememberToken string    `json:"remember_token"`
	LoginType     string    `json:"login_type"`
	Active        bool      `json:"active"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// UserLogin godoc
type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ForgotPassword -
type ForgotPassword struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

// UserModel type that acts as via for datalayer
type UserModel struct{}
