package controllers

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	userRepo    models.UserRepository
	videoRepo   models.VideoRepository
	spaceClient *s3.S3
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(userRepo models.UserRepository, videoRepo models.VideoRepository, spaceClient *s3.S3) *BaseHandler {
	return &BaseHandler{
		userRepo:    userRepo,
		videoRepo:   videoRepo,
		spaceClient: spaceClient,
	}
}
