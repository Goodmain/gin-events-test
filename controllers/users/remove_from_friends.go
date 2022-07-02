package users

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/jwtauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) RemoveFromFriends(c *gin.Context) {
	id := c.Param("id")

	var friend models.User

	if result := h.DB.First(&friend, id); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found"})
		return
	}

	userID, ok := jwtauth.GetUserID(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Unable to decode token")
	}

	var user models.User

	if result := h.DB.First(&user, userID); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found"})
		return
	}

	h.DB.Model(&user).Association("Friends").Delete(&friend)

	c.JSON(http.StatusNoContent, nil)
}
