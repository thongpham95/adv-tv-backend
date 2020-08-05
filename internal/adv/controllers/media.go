package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
	"github.com/thongpham95/adv-tv-backend/internal/adv/utils/responsehandler"
)

type uploadVideoSchema struct {
	DeviceID  string          `json:"device_id"`
	MediaFile *multipart.File `json:"media_file"`
}

// UploadVideo uploads video to Spaces
func (h *BaseHandler) UploadVideo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responsehandler.NewHTTPResponse(false, "Method not allowed", nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(200 << 20) // maxMemory 200MB
	deviceID := r.FormValue("device_id")
	file, handler, err := r.FormFile("media_file")
	if err != nil {
		fmt.Println("Error retrieving file from form-data: ", err)
		responsehandler.NewHTTPResponse(false, "Error retrieving file from form-data: "+err.Error(), nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file %v\n", handler.Filename)
	fmt.Printf("Files size %v\n", handler.Size)
	fmt.Printf("MIME header %v\n", handler.Header)
	advBucket := "adv-storage"
	advKey := deviceID + "/" + handler.Filename // folder name deviceID

	object := s3.PutObjectInput{
		Bucket: aws.String(advBucket),
		Key:    aws.String(advKey),
		Body:   file,
		ACL:    aws.String("public-read"),
	}
	_, putObjectErr := h.spaceClient.PutObject(&object)
	if putObjectErr != nil {
		fmt.Println(putObjectErr.Error())
		responsehandler.NewHTTPResponse(false, "Error uploading file : "+putObjectErr.Error(), nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}

	// save data in db
	mediaItem := models.MediaItem{
		DeviceID: deviceID,
		Key:      advKey,
	}

	h.mediaRepo.Create(deviceID, mediaItem)

	responsehandler.NewHTTPResponse(true, "File uploaded", nil).SuccessResponse(w)
}
