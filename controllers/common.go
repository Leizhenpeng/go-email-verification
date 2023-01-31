package controllers

import (
	"github.com/gin-gonic/gin"
	"leizhenpeng/go-email-verification/models"
	_ "leizhenpeng/go-email-verification/models"
)

type errorString struct {
	string
}

// Ping godoc
//	@Summary	Test if server is alive
//	@Schemes
//	@Tags		Health
//	@Produce	json
// @Success 200 object models.Resp 返回列表
// @Failure 500 object models.Resp 查询失败
//	@Router		/ping [get]
func Ping(c *gin.Context) {
	models.OkWithMessage("pong", c)
}
