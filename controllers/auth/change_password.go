package auth

import (
	"events-hackathon-go/core/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChangePasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func (h handler) ChangePassword(c *gin.Context) {
	data := ChangePasswordRequest{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return

	}

	var user models.User

	if result := h.DB.Where(&models.User{ResetPasswordToken: data.Token}).First(&user); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Incorrect token"})
		return
	}

	user.RawPassword = data.NewPassword
	user.ResetPasswordToken = ""

	if result := h.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Unable update password"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
