package jwtauth

import (
	"events-hackathon-go/core/models"
	"os"
	"strconv"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

func GenerateToken(user models.User) (string, error) {
	secret := os.Getenv("TOKEN_SECRET")
	claims := jwt_lib.StandardClaims{
		Id:        strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func DecodeToken(tokenString string) (string, bool) {
	secret := os.Getenv("TOKEN_SECRET")
	token, _ := jwt_lib.ParseWithClaims(tokenString, &jwt_lib.StandardClaims{}, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*jwt_lib.StandardClaims); ok && token.Valid {
		return claims.Id, true
	} else {
		return "", false
	}
}
