package events

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/jwtauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) AssignOnEvent(c *gin.Context) {
	id := c.Param("id")

	var event models.Event

	if result := h.DB.First(&event, id); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Event not found"})
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

	h.DB.Model(&event).Association("Users").Append(&user)

	c.JSON(http.StatusNoContent, nil)
}
