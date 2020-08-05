package controllers

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
)

// UserFromReqBody ..
type UserFromReqBody struct {
	ID string `json:"user_id"`
}

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	userRepo    models.UserRepository
	mediaRepo   models.MediaItemRepository
	spaceClient *s3.S3
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(userRepo models.UserRepository, mediaRepo models.MediaItemRepository, spaceClient *s3.S3) *BaseHandler {
	return &BaseHandler{
		userRepo:    userRepo,
		mediaRepo:   mediaRepo,
		spaceClient: spaceClient,
	}
}

// GetUserIDFromHeader return user id from request header
func (h *BaseHandler) GetUserIDFromHeader(r *http.Request) string {
	return r.Header["Userid"][0]
}
