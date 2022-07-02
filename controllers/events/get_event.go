package events

import (
	"events-hackathon-go/core/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetEvent(c *gin.Context) {
	id := c.Param("id")

	var event models.Event

	if result := h.DB.Preload("Users").First(&event, id); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}
