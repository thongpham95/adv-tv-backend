package models

// Video model
type Video struct {
	ID           string `json:"id"`
	DeviceID     string `json:"device_id"`
	Title        string `json:"title"`
	Link         string `json:"link"`
	DownloadLink string `json:"download_link"`
}

// VideoRepository exported
type VideoRepository interface {
	Upload(deviceID string, video Video) error
}
