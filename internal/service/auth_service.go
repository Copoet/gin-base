package services

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("your_secret_key")

func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
