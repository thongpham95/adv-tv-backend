package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/thongpham95/adv-tv-backend/internal/adv/controllers"
	"github.com/thongpham95/adv-tv-backend/internal/adv/middlewares"
	postgres "github.com/thongpham95/adv-tv-backend/internal/adv/postgres"
	"github.com/thongpham95/adv-tv-backend/internal/adv/repositories"
	spaces "github.com/thongpham95/adv-tv-backend/internal/adv/spaces"
	"github.com/thongpham95/adv-tv-backend/internal/adv/utils/errorhandler"
)

// ClientError is an error whose details to be shared with client.

// Use as a wrapper around the handler functions.
type rootHandler func(http.ResponseWriter, *http.Request) error

func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function
	if err == nil {
		return
	}
	// This is where our error handling logic starts.
	log.Printf("An error occured: %v", err) // Log the error.

	errorhandler.HTTPErrorResponse(w, err)
}

// Done facilitating error handler

func handleRequests(db *sql.DB, spaceClient *s3.S3) {
	userRepo := repositories.NewUserRepo(db)
	videoRepo := repositories.NewVideoRepo(db)
	controllerHandler := controllers.NewBaseHandler(userRepo, videoRepo, spaceClient)
	// http.Handle("/login", middlewares.IsAuthorized(rootHandler(controllerHandler.Login)))
	http.Handle("/login", rootHandler(controllerHandler.Login))
	http.Handle("/video/upload", middlewares.IsAuthorized(rootHandler(controllerHandler.UploadVideo)))
}

func main() {
	// Spinning db
	db := postgres.ConnectDB()
	// End of Spinning db
	// DigitlOcean Space client spinning up
	spaceClient := spaces.NewSpaceClient()

	fmt.Println("This is a product with luv from lovely Gophers! It is running on port 9000")

	handleRequests(db, spaceClient)

	// log.Fatal is exported from "net/http"
	log.Fatal(http.ListenAndServe(":9000", nil))
}
