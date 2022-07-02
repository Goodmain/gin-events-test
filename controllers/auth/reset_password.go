package auth

import (
	"events-hackathon-go/core/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (h handler) ResetPassword(c *gin.Context) {
	data := ResetPasswordRequest{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return

	}

	var user models.User

	if result := h.DB.Where(&models.User{Email: data.Email}).First(&user); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Incorrect email"})
		return
	}

	user.GenerateResetPasswordToken()

	if result := h.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Unable save token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reset_password_token": user.ResetPasswordToken})
}
