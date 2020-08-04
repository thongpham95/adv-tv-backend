package repositories

import (
	"database/sql"
	"fmt"

	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
)

// VideoRepo implements models.UserRepository
type VideoRepo struct {
	db *sql.DB
}

// NewVideoRepo ..
func NewVideoRepo(db *sql.DB) *VideoRepo {
	return &VideoRepo{
		db: db,
	}
}

// Upload upload video to Spaces
func (r *VideoRepo) Upload(deviceID string, video models.Video) error {
	fmt.Println("Uploading video to ", deviceID)
	return nil
}
