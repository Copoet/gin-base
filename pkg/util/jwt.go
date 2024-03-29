package util

import (
	"errors"
	"gin-base/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var SecretKey = []byte(config.Config.JwtConfig.Secret)

// 生成token
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

// 验证token
func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 判断token是否过期
func IsTokenValid(tokenString string) bool {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	return err == nil
}
