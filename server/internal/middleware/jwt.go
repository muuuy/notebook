package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Middlware to generate a token
func GenerateToken(username string) (string, error) {

	// Create new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expiry
	})
}
