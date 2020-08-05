package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/thongpham95/adv-tv-backend/internal/adv/models"
	advJWT "github.com/thongpham95/adv-tv-backend/internal/adv/utils/jwt"
	"github.com/thongpham95/adv-tv-backend/internal/adv/utils/responsehandler"
	"golang.org/x/crypto/bcrypt"
)

type loginAndSignUpSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login exported
func (h *BaseHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		responsehandler.NewHTTPResponse(false, "Method not allowed", nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body) // Read request body
	if err != nil {
		responsehandler.NewHTTPResponse(false, "Request body read error : "+err.Error(), nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}

	// Parse body as json
	var schema loginAndSignUpSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		responsehandler.NewHTTPResponse(false, "Bad request : invalid JSON, "+err.Error(), nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}

	fmt.Println("Find account in db...")
	user, err := h.userRepo.FindByEmail(schema.Email)
	if err != nil {
		responsehandler.NewHTTPResponse(false, "No user found", nil).ErrorResponse(w, http.StatusNotFound)
		return
	}
	// Verify password
	verifyErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(schema.Password))
	if verifyErr != nil {
		log.Println("Verify password error: ", verifyErr)
		responsehandler.NewHTTPResponse(false, "Wrong password", nil).ErrorResponse(w, http.StatusForbidden)
		return
	}
	// Verify password*
	fmt.Println("Generating token...")
	tokenString, err := advJWT.GenerateJWT(user.ID)
	if err != nil {
		log.Fatalln("Error generating token string")
	}
	token := advJWT.Token{
		Token: tokenString,
	}
	responsehandler.NewHTTPResponse(true, "Login successfully", token).SuccessResponse(w)
}

// SignUp exported
func (h *BaseHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.GetContenTypeFromHeader(r)
	if r.Method != http.MethodPost {
		responsehandler.NewHTTPResponse(false, "Method not allowed", nil).ErrorResponse(w, http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body) // Read request body
	if err != nil {
		responsehandler.NewHTTPResponse(false, "Request body read error : "+err.Error(), nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}

	// Parse body as json
	var schema loginAndSignUpSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		responsehandler.NewHTTPResponse(false, "Bad request : invalid JSON, "+err.Error(), nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}

	// Verify user here
	fmt.Println("Find account in db...")
	foundUser, err := h.userRepo.FindByEmail(schema.Email)
	if foundUser != nil && err == nil {
		responsehandler.NewHTTPResponse(false, "User already exist", nil).ErrorResponse(w, http.StatusConflict)
		return
	}
	// Verify user here*

	// Verify email here
	// -----------------
	// Verify email here*

	// Create account here
	// hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(schema.Password), bcrypt.MinCost)
	if err != nil {
		responsehandler.NewHTTPResponse(false, "Cannot encode user password"+err.Error(), nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}
	user := &models.User{Email: schema.Email, Password: string(hash)}
	// Save in database
	saveUserErr := h.userRepo.Create(user)
	if saveUserErr != nil {
		responsehandler.NewHTTPResponse(false, "Error saving user into db:"+saveUserErr.Error(), nil).ErrorResponse(w, http.StatusInternalServerError)
		return
	}
	// Create account here*
	responsehandler.NewHTTPResponse(true, "Sign up successfully", nil).SuccessResponse(w)
}
