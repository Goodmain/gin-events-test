package events

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/ticketmaster"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handler) SearchEvents(c *gin.Context) {
	city := c.Query("city")

	if strings.Trim(city, " ") == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "City is empty"})
		return
	}

	var events []models.Event

	if result := h.DB.Where(&models.Event{City: city}).Find(&events); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during request"})
		return
	}

	if len(events) == 0 {
		if events, ok := ticketmaster.LoadEvents(city); ok && len(events) > 0 {
			if result := h.DB.Create(&events); result.Error != nil {
				c.JSON(http.StatusInternalServerError, result.Error)
				return
			}

			c.JSON(http.StatusOK, gin.H{"events": events})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}
