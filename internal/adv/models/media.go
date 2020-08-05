package models

// MediaItem model
type MediaItem struct {
	ID       string `json:"id"`
	DeviceID string `json:"device_id"`
	IsVideo  bool   `json:"is_video"`
	Key      string `json:"key"`
}

// MediaItemRepository exported
type MediaItemRepository interface {
	GetMediaBasedOnID(ID string) (*MediaItem, error)
	GetMediaBasedOnKey(key string) (*MediaItem, error)
	Create(deviceID string, mediaItem *MediaItem) error
}
