package main

import (
	"context"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"leizhenpeng/go-email-verification/initialize"
	"leizhenpeng/go-email-verification/middles"
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
	ctx = context.Background()
	//config load
	initialize.InitConfig(".")
	//mongodb client
	initialize.InitClient(ctx)
	defer initialize.CloseClient(ctx)
	//di services
	services.InitUserCollection()
	userService := services.NewUserServicesImpl(services.GetUserCollection(), ctx)
	//gin
	server := gin.Default()
	server.Use(requestid.New())
	server.Use(middles.AddCors())

	api := server.Group("/api")
	routers.InitCommonRouter(api)
	routers.InitUserRouter(ctx, userService, api)
	routers.InitSwaggerRouter(server)

	//swagger url
	log.Println("swagger url:" + initialize.GetConfig().BaseUrl + "/swagger/index.html")
	log.Fatal(server.Run(":" + initialize.GetConfig().Port))
}
