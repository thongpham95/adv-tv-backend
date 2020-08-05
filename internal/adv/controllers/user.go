package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	advJWT "github.com/thongpham95/adv-tv-backend/internal/adv/utils/jwt"
	"github.com/thongpham95/adv-tv-backend/internal/adv/utils/responsehandler"
)

type loginSchema struct {
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
	var schema loginSchema
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
	fmt.Println("Logging user:", user)
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
