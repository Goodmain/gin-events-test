package events

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/ticketmaster"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SearchEventsFilters struct {
	models.SearchFilters
	City string `form:"city" json:"city"`
}

func (s *SearchEventsFilters) Filters(queryField string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.City != "" {
			return s.SearchFilters.Filters(queryField)(db.Where("city = ?", s.City))
		}

		return s.SearchFilters.Filters(queryField)(db)
	}
}

func (h handler) SearchEvents(c *gin.Context) {
	filters := SearchEventsFilters{}

	if err := c.Bind(&filters); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error})
		return
	}

	var events []models.Event

	if result := h.DB.Scopes(filters.Filters("name")).Scopes(filters.Paginate()).Find(&events); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during request"})
		return
	}

	if len(events) == 0 && filters.City != "" { // TODO: move data crawling to the repository
		if events, ok := ticketmaster.LoadEvents(filters.City); ok && len(events) > 0 {
			if result := h.DB.Create(&events); result.Error != nil {
				c.JSON(http.StatusInternalServerError, result.Error)
				return
			}

			if result := h.DB.Scopes(filters.Filters("name")).Scopes(filters.Paginate()).Find(&events); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during request"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"events": events})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}
