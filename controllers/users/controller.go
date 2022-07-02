package users

import (
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

	routes := r.Group("/users")
	routes.GET("/profile", h.GetProfile)
	routes.PUT("/profile", h.UpdateProfile)
	routes.PUT("/change-password", h.ChangePassword)
}
