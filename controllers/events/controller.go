package events

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

	routes := r.Group("/events")
	routes.Use(jwt.Auth(secret))

	routes.GET("/", h.SearchEvents)
	routes.GET("/:id", h.GetEvent)
	routes.PUT("/:id/assign", h.AssignOnEvent)
	routes.PUT("/:id/unassign", h.UnassignFromEvent)
}
