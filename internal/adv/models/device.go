package models

// Device model
type Device struct {
	ID             string `json:"id"`
	Owner          string `json:"owner"`
	Name           string `json:"name"`
	Model          string `json:"model"`
	Serial         string `json:"serial"`
	Manufacturer   string `json:"manufacturer"`
	AppVersion     string `json:"app_version"`
	AndroidVersion string `json:"android_version"`
	LastOpen       string `json:"last_open"`
	LastUpdate     string `json:"last_update"`
}

// DeviceRepository contract bewteen controller layer and repo layer
type DeviceRepository interface {
	Get(deviceID string) (*Device, error)
	Create(userID string, device *Device) error
}
