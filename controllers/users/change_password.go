package users

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/jwtauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func (h handler) ChangePassword(c *gin.Context) {
	data := ChangePasswordRequest{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, ok := jwtauth.GetUserID(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Unable to decode token")
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found"})
		return
	}

	if !user.IsValidPassword(data.Password) {
		c.JSON(http.StatusUnprocessableEntity, "Incorrect password")
	}

	user.RawPassword = data.NewPassword

	if result := h.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Unable to update profile"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
