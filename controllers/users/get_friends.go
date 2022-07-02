package users

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/jwtauth"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SearchFriendsFilters struct {
	models.SearchFilters
	Email string `form:"email" json:"email"`
}

func (s *SearchFriendsFilters) Filters(queryField string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Email != "" {
			return s.SearchFilters.Filters(queryField)(db.Where("email = ?", s.Email))
		}

		return s.SearchFilters.Filters(queryField)(db)
	}
}

func (h handler) GetFriends(c *gin.Context) {
	id, ok := jwtauth.GetUserID(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Unable to decode token")
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found"})
		return
	}

	filters := SearchFriendsFilters{}

	if err := c.Bind(&filters); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error})
		return
	}

	var friends []models.User

	if result := h.DB.Model(&user).Scopes(filters.Filters("name")).Scopes(filters.Paginate()).Association("Friends").Find(&friends); result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"friends": friends})
}
