package users

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

	routes := r.Group("/users")
	routes.Use(jwt.Auth(secret))

	routes.GET("/profile", h.GetProfile)
	routes.PUT("/profile", h.UpdateProfile)
	routes.PUT("/change-password", h.ChangePassword)
}
