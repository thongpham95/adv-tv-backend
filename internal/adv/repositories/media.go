package repositories

import (
	"database/sql"
	"log"

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

// GetMediaBasedOnID return media based on its ID
func (r *MediaItem) GetMediaBasedOnID(ID string) (*models.MediaItem, error) {
	var media models.MediaItem
	err := r.db.QueryRow("SELECT * FROM media_item WHERE id = $1 LIMIT 1", ID).Scan(&media.ID, &media.DeviceID, &media.IsVideo, &media.Key)
	if err != nil {
		log.Println("Failed to execute query: ", err)
		return nil, err
	}
	return &media, nil
}

// Create upload media item to Spaces
func (r *MediaItem) Create(deviceID string, mediaItem *models.MediaItem) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare("INSERT INTO media_item(device_id, isvideo, key) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	if _, err := stmt.Exec(deviceID, mediaItem.IsVideo, mediaItem.Key); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// GetMediaBasedOnKey return media based on its key
func (r *MediaItem) GetMediaBasedOnKey(key string) (*models.MediaItem, error) {
	var media models.MediaItem
	err := r.db.QueryRow("SELECT * FROM media_item WHERE key = $1", key).Scan(&media.ID, &media.DeviceID, &media.IsVideo, &media.Key)
	if err != nil {
		log.Println("Failed to execute query: ", err)
		return nil, err
	}
	return &media, nil
}
