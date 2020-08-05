package repositories

import (
	"database/sql"
	"log"

	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
)

// Device implements models.UserRepository
type Device struct {
	db *sql.DB
}

// NewDevice ..
func NewDevice(db *sql.DB) *Device {
	return &Device{
		db: db,
	}
}

// Get return device info based on device ID
func (r *Device) Get(deviceID string) (*models.Device, error) {
	var device models.Device
	err := r.db.QueryRow("SELECT * FROM device WHERE id = $1", deviceID).Scan(&device.ID, &device.Owner, &device.Name, &device.Model, &device.Serial, &device.Manufacturer, &device.AppVersion, &device.AndroidVersion, &device.LastOpen, &device.LastUpdate)
	if err != nil {
		log.Println("Failed to execute query: ", err)
		return nil, err
	}
	return &device, nil
}

// Create upload media item to Spaces
func (r *Device) Create(userID string, device *models.Device) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare("INSERT INTO device(owner, name, model, serial, manufacturer, app_version, android_version, last_open, last_update) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		return err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	if _, err := stmt.Exec(userID, device.Name, device.Model, device.Serial, device.Manufacturer, device.AppVersion, device.AndroidVersion, device.LastOpen, device.LastUpdate); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
