package auth

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/jwtauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) RefreshToken(c *gin.Context) {
	id, ok := jwtauth.GetUserID(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Unable to decode token")
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found"})
		return
	}

	token, err := jwtauth.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to generate token")
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
