package auth

import (
	"os"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	secret := os.Getenv("TOKEN_SECRET")

	routes := r.Group("/auth")
	routes.POST("/login", h.Login)
	routes.POST("/register", h.Register)

	routes.Use(jwt.Auth(secret))
	routes.POST("/refresh-token", h.RefreshToken)
}
