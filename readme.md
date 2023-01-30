## go-email-verification




## Features


- 🐰 数据持久化 MongoDB
- 🦊 验证码缓存 Redis
- 🐼 邮件发送 SMTP
- 🐮 邮件模板 HTML
- 🐦 接口文档 Swagger


## SHOT

- SWAG
``` golang
type PingResponse struct {
	CommonResponse
	Msg string `json:"Message" example:"pong"`
}

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

```

## 相关阅读

[Golang and MongoDB using the official mongo-driver](https://wb.id.au/computer/golang-and-mongodb-using-the-mongo-go-driver/)

[Gin middleware with Swagger 2.0](https://github.com/swaggo/gin-swagger)

[使用swaggo自动生成Restful API文档](https://ieevee.com/tech/2018/04/19/go-swag.html)
