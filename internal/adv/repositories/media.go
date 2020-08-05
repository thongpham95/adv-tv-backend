package repositories

import (
	"database/sql"
	"fmt"

	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
)

// MediaItem implements models.UserRepository
type MediaItem struct {
	db *sql.DB
}

// NewMediaItem ..
func NewMediaItem(db *sql.DB) *MediaItem {
	return &MediaItem{
		db: db,
	}
}

// Create upload media item to Spaces
func (r *MediaItem) Create(deviceID string, mediaItem models.MediaItem) error {
	fmt.Println("Uploading video to", deviceID)
	return nil
}
