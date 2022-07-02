package users

import (
	"events-hackathon-go/core/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
