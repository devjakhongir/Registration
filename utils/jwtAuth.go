package utils

import (
	"time"
	"app/models"
	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(jwtKey string) (string, error) {

	expirationTime := time.Now().Add(time.Minute * 25)

	claims := &models.Claims {
		StandardClaims:jwt.StandardClaims {
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))

	return tokenString, err
}