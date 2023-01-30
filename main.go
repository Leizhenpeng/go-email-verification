package main

import (
	"github.com/gin-gonic/gin"
	"leizhenpeng/go-email-verification/config"
	"leizhenpeng/go-email-verification/middles"
	"leizhenpeng/go-email-verification/models"
	"leizhenpeng/go-email-verification/routers"
	"log"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config: ", err)
		return
	}
	models.InitClient()
	defer models.CloseClient()
	server := gin.Default()
	server.Use(middles.AddCors())

	api := server.Group("/api")
	routers.CommonRoute(api)
	routers.SwaggerRoute(server)

	log.Fatal(server.Run(":" + config.GetConfig().Port))
}
