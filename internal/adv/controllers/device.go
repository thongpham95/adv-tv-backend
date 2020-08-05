package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lib/pq"
	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
	"github.com/thongpham95/adv-tv-backend/internal/adv/utils/responsehandler"
)

type addDeviceSchema struct {
	Name           string `json:"name"`
	Model          string `json:"model"`
	Serial         string `json:"serial"`
	Manufacturer   string `json:"manufacturer"`
	AppVersion     string `json:"app_version"`
	AndroidVersion string `json:"android_version"`
}

// AddDevice add new device to user
func (h *BaseHandler) AddDevice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responsehandler.NewHTTPResponse(false, "Method not allowed", nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}
	userID := h.GetUserIDFromHeader(r)
	fmt.Println("userID: ", userID)
	var schema addDeviceSchema
	if err := h.ParseNormalRequestBody(r, &schema); err != nil {
		log.Print(err)
		responsehandler.NewHTTPResponse(false, "Bad request", nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}
	now := time.Now() // current local time
	device := &models.Device{
		Name:           schema.Name,
		Model:          schema.Model,
		Serial:         schema.Serial,
		Manufacturer:   schema.Manufacturer,
		AppVersion:     schema.AppVersion,
		AndroidVersion: schema.AndroidVersion,
		LastOpen:       string(pq.FormatTimestamp(now)),
		LastUpdate:     string(pq.FormatTimestamp(now)),
	}
	if err := h.deviceRepo.Create(userID, device); err != nil {
		log.Print(err)
		responsehandler.NewHTTPResponse(false, "Error adding device to user", nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}
	responsehandler.NewHTTPResponse(true, "Device added", nil).SuccessResponse(w)
}
