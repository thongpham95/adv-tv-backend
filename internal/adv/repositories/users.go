package repositories

import (
	"database/sql"
	"log"

	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo ..
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// FindByEmail ..
func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT * FROM adv.user WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
		return nil, err
	}
	return &user, nil
}

// Save ..
func (r *UserRepo) Save(user *models.User) error {
	return nil
}
