package controllers

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
	"github.com/thongpham95/adv-tv-backend/internal/adv/utils/responsehandler"
)

type uploadVideoSchema struct {
	DeviceID  string          `json:"device_id"`
	MediaFile *multipart.File `json:"media_file"`
}

// GetMediaURL return pre-signed URL of a requested file
func (h *BaseHandler) GetMediaURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responsehandler.NewHTTPResponse(false, "Method not allowed", nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}
	mediaID := r.URL.Query().Get("media_id")
	if len(mediaID) < 1 {
		responsehandler.NewHTTPResponse(false, "Missing media ID", nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}
	// get media key
	media, err := h.mediaRepo.GetMediaBasedOnID(mediaID)
	if media == nil || err != nil || len(media.Key) < 1 {
		log.Println(err)
		responsehandler.NewHTTPResponse(false, "File not found", nil).ErrorResponse(w, http.StatusNotFound)
		return
	}
	advBucket := "adv-storage"
	mediaKey := media.Key

	// get file based on media key
	req, _ := h.spaceClient.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(advBucket),
		Key:    aws.String(mediaKey),
	})

	// get URL which last in 1 day
	urlStr, err := req.Presign(24 * time.Hour)
	if err != nil {
		log.Println(err)
		responsehandler.NewHTTPResponse(false, "Error getting file from storage service", nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}
	type customResponse struct {
		URL string `json:"file_url"`
	}
	customeRes := customResponse{
		URL: urlStr,
	}
	responsehandler.NewHTTPResponse(true, "Get file url successfully", customeRes).SuccessResponse(w)
}

// UploadVideo uploads video to Spaces
func (h *BaseHandler) UploadVideo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		responsehandler.NewHTTPResponse(false, "Method not allowed", nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}

	// Parse Multipart Form in order to retrieve file data
	r.ParseMultipartForm(200 << 20) // maxMemory 200MB
	deviceID := r.FormValue("device_id")
	file, handler, err := r.FormFile("media_file")
	if err != nil {
		fmt.Println("Error retrieving file from form-data: ", err)
		responsehandler.NewHTTPResponse(false, "Error retrieving file from form-data: "+err.Error(), nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file: %v\n", handler.Filename)
	fmt.Printf("Files size: %v\n", handler.Size)
	fmt.Printf("MIME header: %v\n", handler.Header["Content-Type"][0])
	advBucket := "adv-storage"
	mediaKey := deviceID + "/" + handler.Filename // folder name deviceID

	// check whether device already uploaded the file
	if media, _ := h.mediaRepo.GetMediaBasedOnKey(mediaKey); media != nil {
		responsehandler.NewHTTPResponse(false, "Device already uploaded this file", nil).ErrorResponse(w, http.StatusConflict)
		return
	}
	object := s3.PutObjectInput{
		Bucket: aws.String(advBucket),
		Key:    aws.String(mediaKey),
		Body:   file,
		ACL:    aws.String("private"),
	}
	_, putObjectErr := h.spaceClient.PutObject(&object)
	if putObjectErr != nil {
		log.Println(putObjectErr)
		responsehandler.NewHTTPResponse(false, "Error uploading file", nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}

	// save data in db
	fileType := strings.Split(handler.Header["Content-Type"][0], "/")[0]
	mediaItem := &models.MediaItem{
		IsVideo: !(fileType == "image"),
		Key:     mediaKey,
	}

	// check whether device exist
	if device, err := h.deviceRepo.Get(deviceID); device == nil || err != nil {
		responsehandler.NewHTTPResponse(false, "Device not found", nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}

	// add new media item to database
	if err := h.mediaRepo.Create(deviceID, mediaItem); err != nil {
		log.Println(err)
		responsehandler.NewHTTPResponse(false, "Error uploading file", nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}
	responsehandler.NewHTTPResponse(true, "File uploaded successfully", nil).SuccessResponse(w)
}
