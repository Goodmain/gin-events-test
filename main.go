package main

import (
	"events-hackathon-go/controllers/auth"
	"events-hackathon-go/core/services/pg"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("./configs/server.json")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	h := pg.Init()

	//users.RegisterRoutes(r, h)
	//events.RegisterRoutes(r, h)
	auth.RegisterRoutes(r, h)

	r.Run(port)
}
