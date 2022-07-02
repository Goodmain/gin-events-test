package auth

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/jwtauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Login(c *gin.Context) {
	data := LoginRequest{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return

	}

	var user models.User

	if result := h.DB.Where(&models.User{Email: data.Email}).First(&user); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Incorrect email or password"})
		return
	}

	if !user.IsValidPassword(data.Password) {
		c.JSON(http.StatusUnprocessableEntity, "Incorrect password")
	}

	token, err := jwtauth.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to generate token")
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
