package models

// MediaItem model
type MediaItem struct {
	ID       string `json:"id"`
	DeviceID string `json:"device_id"`
	Key      string `json:"key"`
}

// MediaItemRepository exported
type MediaItemRepository interface {
	Create(deviceID string, mediaItem MediaItem) error
}
