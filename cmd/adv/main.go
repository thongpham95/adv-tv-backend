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
)

// ClientError is an error whose details to be shared with client.

// Done facilitating error handler

func handleRequests(db *sql.DB, spaceClient *s3.S3) {
	userRepo := repositories.NewUserRepo(db)
	mediaRepo := repositories.NewMediaItem(db)
	deviceRepo := repositories.NewDevice(db)
	controllerHandler := controllers.NewBaseHandler(userRepo, mediaRepo, deviceRepo, spaceClient)
	http.HandleFunc("/login", controllerHandler.Login)                                     // POST
	http.HandleFunc("/signup", controllerHandler.SignUp)                                   // POST
	http.Handle("/device", middlewares.IsAuthorized(controllerHandler.AddDevice))          // POST
	http.Handle("/media/upload", middlewares.IsAuthorized(controllerHandler.UploadVideo))  // POST
	http.Handle("/media", middlewares.IsAuthorized(controllerHandler.GetMediaURL))         // GET
	http.Handle("/app/check", middlewares.IsAuthorized(controllerHandler.CheckAppVersion)) // GET
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
