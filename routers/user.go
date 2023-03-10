package routers

import (
	"context"
	"github.com/gin-gonic/gin"
	"leizhenpeng/go-email-verification/controllers"
	"leizhenpeng/go-email-verification/middles"
	"leizhenpeng/go-email-verification/services"
)

func InitUserRouter(ctx context.Context, userService services.UserServices, router *gin.RouterGroup) {
	controllersImpl := controllers.NewUserControllersImpl(ctx, userService)
	router.POST("/signup", controllersImpl.SignUp)
	authMiddleware, err := middles.InitAuthMiddlewares(controllersImpl)
	if err != nil {
		panic(err)
	}
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/refresh_token", authMiddleware.RefreshHandler)
	router.GET("/verify_email", controllersImpl.VerifyEmail)

	{
		auth := router.Use(authMiddleware.MiddlewareFunc())
		auth.POST("/logout", authMiddleware.LogoutHandler)
		auth.GET("/me", controllersImpl.Info)
		auth.POST("/send_email", controllersImpl.SendEmail)
	}
}
