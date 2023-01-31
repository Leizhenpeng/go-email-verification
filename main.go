package main

import (
	"context"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"leizhenpeng/go-email-verification/config"
	"leizhenpeng/go-email-verification/middles"
	"leizhenpeng/go-email-verification/models"
	"leizhenpeng/go-email-verification/routers"
	"leizhenpeng/go-email-verification/services"
	"log"
)

var ctx context.Context

// @title Email Verification API
// @version 1.0.0
// @description Create and verify email addresses
// @host 127.0.0.1:8000
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config: ", err)
		return
	}
	ctx = context.Background()
	models.InitClient(ctx)
	defer models.CloseClient(ctx)

	models.InitUserCollection()
	userService := services.NewUserServicesImpl(models.GetUserCollection(), ctx)

	server := gin.Default()
	server.Use(requestid.New())
	server.Use(middles.AddCors())

	routers.SwaggerRoute(server)

	api := server.Group("/api")
	routers.CommonRoute(api)

	routers.InitUserRouter(ctx, userService, api)

	log.Fatal(server.Run(":" + config.GetConfig().Port))
}
