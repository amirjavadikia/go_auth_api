package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtSecret = []byte("4DYodcKIOUfYCATnvBVqOXrLwyWEPvvAE8/t6WHS56A=")

func GenerateJwt(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	clamis := &jwt.RegisteredClaims{
		Subject:   string(userID),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
