package auth

import (
	"events-hackathon-go/core/models"
	"events-hackathon-go/core/services/jwtauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Register(c *gin.Context) {
	data := RegisterRequest{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	user.Email = data.Email
	user.Name = data.Name
	user.RawPassword = data.Password

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	token, err := jwtauth.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to generate token")
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
