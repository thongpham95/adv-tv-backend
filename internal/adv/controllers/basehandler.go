package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
	advcontext "github.com/thongpham95/adv-tv-backend/internal/adv/utils/context"
)

// UserFromReqBody ..
type UserFromReqBody struct {
	ID string `json:"user_id"`
}

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	userRepo    models.UserRepository
	mediaRepo   models.MediaItemRepository
	deviceRepo  models.DeviceRepository
	spaceClient *s3.S3
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(userRepo models.UserRepository, mediaRepo models.MediaItemRepository, deviceRepo models.DeviceRepository, spaceClient *s3.S3) *BaseHandler {
	return &BaseHandler{
		userRepo:    userRepo,
		mediaRepo:   mediaRepo,
		deviceRepo:  deviceRepo,
		spaceClient: spaceClient,
	}
}

// ParseNormalRequestBody CANNOT parse multipart request, schema must be passed as a pointer value
func (h *BaseHandler) ParseNormalRequestBody(r *http.Request, schema interface{}) error {
	body, err := ioutil.ReadAll(r.Body) // Read request body
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, schema); err != nil {
		return err
	}
	return nil
}

// GetUserIDFromContext return user id from request context
func (h *BaseHandler) GetUserIDFromContext(r *http.Request) string {
	if ctxUserID, ok := r.Context().Value(advcontext.CtxKey("auth-token")).(string); ok {
		return ctxUserID
	}
	return ""
}

// GetContenTypeFromHeader return Content-Type from request header
func (h *BaseHandler) GetContenTypeFromHeader(r *http.Request) {
	fmt.Println("Content-Type:", r.Header["Content-Type"][0])
}
