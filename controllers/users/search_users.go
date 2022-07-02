package users

import (
	"events-hackathon-go/core/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SearchUsersFilters struct {
	models.SearchFilters
	Email string `form:"email" json:"email"`
}

func (s *SearchUsersFilters) Filters(queryField string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Email != "" {
			return s.SearchFilters.Filters(queryField)(db.Where("email = ?", s.Email))
		}

		return s.SearchFilters.Filters(queryField)(db)
	}
}

func (h handler) SearchUsers(c *gin.Context) {
	filters := SearchUsersFilters{}

	if err := c.Bind(&filters); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error})
		return
	}

	var users []models.User

	if result := h.DB.Scopes(filters.Filters("name")).Scopes(filters.Paginate()).Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
