package middlewares

import (
	"fmt"
	"net/http"

	errorHandler "github.com/thongpham95/adv-tv-backend/internal/adv/utils/errorhandler"
	advjwt "github.com/thongpham95/adv-tv-backend/internal/adv/utils/jwt"
)

// IsAuthorized exported
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			fmt.Println("Unauthorized")
			errorHandler.HTTPErrorResponse(w, errorHandler.NewHTTPError(nil, http.StatusUnauthorized, "Unauthorized"))
		} else {
			fmt.Println("Getting token from header")
			ok, err := advjwt.ValidateToken(r.Header["Authorization"][0])
			if err != nil {
				fmt.Println("Err Getting token from header: ", err)
				errorHandler.HTTPErrorResponse(w, errorHandler.NewHTTPError(err, http.StatusUnauthorized, "Unauthorized"))
			}
			if ok == true {
				fmt.Println("Token valid: ", ok)
				endpoint(w, r)
			}
		}
	})
}
