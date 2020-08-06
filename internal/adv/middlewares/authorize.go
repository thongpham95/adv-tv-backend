package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"

	advcontext "github.com/thongpham95/adv-tv-backend/internal/adv/utils/context"
	advjwt "github.com/thongpham95/adv-tv-backend/internal/adv/utils/jwt"
	responsehandler "github.com/thongpham95/adv-tv-backend/internal/adv/utils/responsehandler"
)

// UserBody ..
type UserBody struct {
	ID    string `json:"user_id"`
	Email string `json:"user_email"`
}

// IsAuthorized exported
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			responsehandler.NewHTTPResponse(false, "Unauthorized", nil).ErrorResponse(w, http.StatusUnauthorized)
		} else {
			tokenStr, err := advjwt.ValidateToken(r.Header["Authorization"][0])
			if err != nil {
				log.Println("Error getting token from header: " + err.Error())
				responsehandler.NewHTTPResponse(false, err.Error(), nil).ErrorResponse(w, http.StatusUnauthorized)
			}
			if tokenStr != nil {
				ctx := context.WithValue(r.Context(), advcontext.CtxKey("auth-token"), fmt.Sprintf("%v", tokenStr))
				endpoint(w, r.WithContext(ctx))
			}
		}
	})
}
