package middles

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"leizhenpeng/go-email-verification/config"
	"leizhenpeng/go-email-verification/controllers"
	"leizhenpeng/go-email-verification/models"
	"net/http"
	"time"
)

func InitAuthMiddlewares(controllers controllers.UserControllers) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		IdentityKey:      "id",
		Realm:            "email-verification",
		SigningAlgorithm: "HS256",
		Key:              []byte(config.GetConfig().JwtKey),
		Timeout:          time.Hour * time.Duration(config.GetConfig().JwtAccessAge),
		MaxRefresh:       time.Hour * time.Duration(config.GetConfig().JwtRefreshAge),
		TokenLookup:      "header: Authorization",
		TokenHeadName:    "Bearer",
		TimeFunc:         time.Now,
		Authenticator:    controllers.Login,
		Authorizator:     authorizedFunc,
		PayloadFunc:      payloadHandle,
		LoginResponse:    loginResponse,
		LogoutResponse:   logoutResponse,
		Unauthorized:     unauthorizedFunc,
		IdentityHandler:  identityHandler,
	})
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return claims["id"]
}
func payloadHandle(data interface{}) jwt.MapClaims {
	fmt.Printf("data: %+v", data)
	return jwt.MapClaims{
		"id":    data.(*models.DBUserResponse).ID,
		"email": data.(*models.DBUserResponse).Email,
	}
}

func unauthorizedFunc(c *gin.Context, code int, message string) {
	c.JSON(http.StatusUnauthorized, models.Resp{
		RequestId: requestid.Get(c),
		Code:      code,
		Msg:       message,
	})
}
func loginResponse(c *gin.Context, code int, token string, expires time.Time) {
	c.JSON(http.StatusOK, models.Resp{
		RequestId: requestid.Get(c),
		Code:      http.StatusOK,
		Data: models.Token{
			Token:   "Bearer " + token,
			Expires: expires,
		},
		Msg: models.CustomError[models.Ok],
	})
}

// logout godoc
// @Summary 退出登录
// @Id logout
// @Tags User
// @version 1.0
// @Accept application/x-json-stream
// @Success 200 object models.Token 返回列表
// @Failure 400 object models.Resp 查询失败
// @Security ApiKeyAuth
// @Router /logout [post]

func logoutResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, models.Resp{
		RequestId: requestid.Get(c),
		Code:      http.StatusOK,
		Data:      models.OkMsg,
		Msg:       models.CustomError[models.Ok],
	})
}

func authorizedFunc(data interface{}, c *gin.Context) bool {
	return true
}

// refresh_token godoc
// @Summary 刷新token
// @Id refresh_token
// @Tags User
// @version 1.0
// @Accept application/x-json-stream
// @Success 200 object models.Token 返回列表
// @Failure 400 object models.Resp 查询失败
// @Security ApiKeyAuth
// @Router /refresh_token [post]
func _(c *gin.Context) {

}
