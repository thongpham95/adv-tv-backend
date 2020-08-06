package models

// MediaItem model
type MediaItem struct {
	ID       string `json:"id"`
	DeviceID string `json:"device_id"`
	IsVideo  bool   `json:"is_video"`
	Key      string `json:"key"`
}

// AppVersion model
type AppVersion struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Link    bool   `json:"link"`
}

// MediaItemRepository exported
type MediaItemRepository interface {
	GetMediaBasedOnID(ID string) (*MediaItem, error)
	GetMediaBasedOnKey(key string) (*MediaItem, error)
	GetCurrentAppVersion() (string, error)
	Create(deviceID string, mediaItem *MediaItem) error
}
