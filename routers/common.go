package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"leizhenpeng/go-email-verification/controllers"
	"leizhenpeng/go-email-verification/docs"
)

func InitCommonRouter(router *gin.RouterGroup) {
	router.GET("/ping", controllers.Ping)
}

func InitSwaggerRouter(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
