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
	err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		log.Println("Failed to execute query: ", err)
		return nil, err
	}
	return &user, nil
}

// Create insert new user into database
func (r *UserRepo) Create(user *models.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare("INSERT INTO users(email, password) VALUES($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	if _, err := stmt.Exec(user.Email, user.Password); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
