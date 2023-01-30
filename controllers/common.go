package controllers

import (
	"github.com/gin-gonic/gin"
	_ "leizhenpeng/go-email-verification/models"
)

// Ping godoc
// @Summary Test if server is alive
// @Schemes
// @Tags Common
// @Produce json
// @Success 200 {object}  models.PingResponse
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
