package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"leizhenpeng/go-email-verification/controllers"
	"leizhenpeng/go-email-verification/docs"
)

func CommonRoute(router *gin.RouterGroup) {
	router.GET("/ping", controllers.Ping)
}

func SwaggerRoute(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Email Verification API"
	docs.SwaggerInfo.Description = "Create and verify email addresses"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
