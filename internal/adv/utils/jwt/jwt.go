package jwt

import (
	"fmt"
	"time"

	"github.com/thongpham95/adv-tv-backend/internal/adv/constants"

	"github.com/dgrijalva/jwt-go"
)

// Token exported
type Token struct {
	Token string `json:"token"`
}

// GenerateJWT i
func GenerateJWT(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix() // token last for 2 minutes
	tokenString, err := token.SignedString(constants.MySecretKey)
	if err != nil {
		errorResult := fmt.Errorf("Something went wrong: %s", err.Error())
		return "", errorResult
	}
	return tokenString, nil
}

// ValidateToken validate the token string
func ValidateToken(tokenString string) (interface{}, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error when validating token")
		}
		return []byte(constants.MySecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Token validation error: %v", err)
	}

	if token.Valid {
		// for key, val := range claims {
		// 	fmt.Printf("Key: %v, value: %v\n", key, val)
		// }
		return claims["user"], nil
	}
	return nil, fmt.Errorf("Invalid token")
}
