package users

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/jwtauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateProfileRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h handler) UpdateProfile(c *gin.Context) {
	data := UpdateProfileRequest{}

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

	user.Name = data.Name

	if result := h.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Unable to update profile"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
