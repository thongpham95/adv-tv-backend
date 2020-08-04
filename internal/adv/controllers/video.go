package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	errorHandler "github.com/thongpham95/adv-tv-backend/internal/adv/utils/errorhandler"
	"github.com/thongpham95/adv-tv-backend/internal/adv/utils/responsehandler"
)

type uploadVideoSchema struct {
	DeviceID string          `json:"device_id"`
	Video    *multipart.File `json:"video"`
}

// UploadVideo uploads video to Spaces
func (h *BaseHandler) UploadVideo(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Is here")
	if r.Method != http.MethodPost {
		return errorHandler.NewHTTPError(nil, 405, "Method not allowed.")
	}
	// fmt.Println("Read request body")
	// body, err := ioutil.ReadAll(r.Body) // Read request body
	// if err != nil {
	// 	fmt.Println("Request body read error")
	// 	return fmt.Errorf("Request body read error : %v", err)
	// }

	// // Parse body as json
	// var schema uploadVideoSchema
	// if err = json.Unmarshal(body, &schema); err != nil {
	// 	return errorHandler.NewHTTPError(err, 400, "Bad request : invalid JSON")
	// }
	// fmt.Println("Parse body as json DONE")
	// business logic here
	// r.ParseMultipartForm(10 << 20)
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
		return fmt.Errorf(putObjectErr.Error())
	}
	fmt.Println("File uploaded")
	responsehandler.NewHTTPResponse(nil).SuccessResponse(w)

	return nil
}
