package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/thongpham95/adv-tv-backend/internal/adv/utils/responsehandler"
)

type uploadVideoSchema struct {
	DeviceID string          `json:"device_id"`
	Video    *multipart.File `json:"video"`
}

// UploadVideo uploads video to Spaces
func (h *BaseHandler) UploadVideo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Is here")
	if r.Method != http.MethodPost {
		responsehandler.NewHTTPResponse(false, "Method not allowed", nil).ErrorResponse(w, http.StatusBadRequest)
	}
	r.ParseMultipartForm(200)

	file, handler, err := r.FormFile("video")
	if err != nil {
		fmt.Println("Error retrieving file from form-data: ", err)
	}
	defer file.Close()
	fmt.Printf("Uploaded file %v\n", handler.Filename)
	fmt.Printf("Files size %v\n", handler.Size)
	fmt.Printf("MIME header %v\n", handler.Header)
	advBucket := "adv-storage"
	advKey := handler.Filename

	object := s3.PutObjectInput{
		Bucket: aws.String(advBucket),
		Key:    aws.String(advKey),
		Body:   file,
		ACL:    aws.String("public-read"),
	}
	_, putObjectErr := h.spaceClient.PutObject(&object)
	if putObjectErr != nil {
		fmt.Println(putObjectErr.Error())
	}
	responsehandler.NewHTTPResponse(true, "File uploaded", nil).SuccessResponse(w)
}
