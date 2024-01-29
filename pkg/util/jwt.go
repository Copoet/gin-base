package util

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.Claims
}

// 密钥
var SecretKey = []byte("wdt112")

// 生成token
func GenerateToken(username, password string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
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
	return nil, err
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
}

// 验证token有效期
func ValidTokenExpress(tokenString string) bool {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return false
	}
	return true
}

// 删除token
func DeleteToken(tokenString string) bool {

	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return false
	}
	return true
}
