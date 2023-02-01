package models

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Resp http请求响应体
type Resp struct {
	RequestId string      `json:"request_id"`     // 请求ID
	Code      int         `json:"code"`           // 错误代码
	Msg       string      `json:"msg,omitempty"`  // 消息提示
	Data      interface{} `json:"data,omitempty"` // 数据内容
}

const (
	Ok                  = 201
	NotOk               = 405
	Unauthorized        = 401
	Forbidden           = 403
	InternalServerError = 500
)

const (
	OkMsg                  = "操作成功"
	NotOkMsg               = "操作失败"
	UnauthorizedMsg        = "登录过期, 需要重新登录"
	LoginCheckErrorMsg     = "用户名或密码错误"
	ForbiddenMsg           = "无权访问该资源, 请联系网站管理员授权"
	InternalServerErrorMsg = "服务器内部错误"
)

var CustomError = map[int]string{
	Ok:                  OkMsg,
	NotOk:               NotOkMsg,
	Unauthorized:        UnauthorizedMsg,
	Forbidden:           ForbiddenMsg,
	InternalServerError: InternalServerErrorMsg,
}

const (
	ERROR   = 1
	SUCCESS = 0
)

var EmptyArray = []interface{}{}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Resp{
		requestid.Get(c),
		code,
		msg,
		data,
	})
}

func OkResult(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func FailResult(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
