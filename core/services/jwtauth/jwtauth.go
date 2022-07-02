package jwtauth

import (
	"events-hackathon-go/core/models"
	"os"
	"strconv"
	"strings"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(user models.User) (string, error) {
	secret := os.Getenv("TOKEN_SECRET")
	claims := jwt_lib.StandardClaims{
		Id:        strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
	}
	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func GetUserID(c *gin.Context) (string, bool) {
	tokenString := strings.Trim(strings.ReplaceAll(c.Request.Header.Get("Authorization"), "Bearer", ""), " ")

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
