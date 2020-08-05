package models

// User exported
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRepository exported
type UserRepository interface {
	FindByEmail(email string) (*User, error)
	Create(user *User) error
}
