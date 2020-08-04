package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	errorHandler "github.com/thongpham95/adv-tv-backend/internal/adv/utils/errorhandler"
	advJWT "github.com/thongpham95/adv-tv-backend/internal/adv/utils/jwt"
	"github.com/thongpham95/adv-tv-backend/internal/adv/utils/responsehandler"
)

type loginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login exported
func (h *BaseHandler) Login(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return errorHandler.NewHTTPError(nil, 405, "Method not allowed.")
	}

	body, err := ioutil.ReadAll(r.Body) // Read request body
	if err != nil {
		return fmt.Errorf("Request body read error : %v", err)
	}

	// Parse body as json
	var schema loginSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		return errorHandler.NewHTTPError(err, 400, "Bad request : invalid JSON")
	}

	fmt.Println("Find account in db...")
	user, err := h.userRepo.FindByEmail(schema.Email)
	if err != nil {
		return errorHandler.NewHTTPError(nil, 404, "No user found")
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
	responsehandler.NewHTTPResponse(token).SuccessResponse(w)

	return nil
}
