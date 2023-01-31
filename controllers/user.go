package controllers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"leizhenpeng/go-email-verification/models"
	"leizhenpeng/go-email-verification/services"
	"strings"
)

type UserSignUpRequest struct {
	Name     string `json:"name" binding:"required" default:"river"`
	Email    string `json:"email" binding:"required" default:"rivers@88.com""`
	Password string `json:"password" binding:"required" default:"123456"`
}

type DBUserResponse struct {
	Name     string `json:"name" bson:"Name"  binding:"required"`
	Email    string `json:"email" bson:"Email"  binding:"required"`
	ID       string `json:"id" bson:"_id"`
	Verified bool   `json:"verified" bson:"Verified"`
}

type UserControllers interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context) (interface{}, error)
	Info(ctx *gin.Context)
}

func NewUserControllersImpl(ctx context.Context, userS services.UserServices) *UserControllersImpl {
	return &UserControllersImpl{
		ctx:   ctx,
		userS: userS,
	}
}

type UserControllersImpl struct {
	ctx   context.Context
	userS services.UserServices
}

// Info godoc
//	@Summary	获取用户信息
//	@Schemes
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Success	200		{}	string "OK"
// @Security ApiKeyAuth
//	@Router		/me [get]
func (u UserControllersImpl) Info(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := claims["id"].(string)
	userInfo, err := u.userS.FindUserById(id)
	if err != nil {
		return
	}
	models.Result(models.SUCCESS, userInfo, "ok", c)
}

// Login godoc
//	@Summary	登录
//	@Schemes
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		user	body		models.LoginInputReq	true	"User"
//	@Success	200		object	 models.Resp   登录成功
//	@Router		/login [post]
func (u UserControllersImpl) Login(ctx *gin.Context) (interface{}, error) {
	var loginInput *models.LoginInputReq
	err := ctx.ShouldBindJSON(&loginInput)
	if err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	dbResult, err := u.userS.Login(loginInput)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return nil, jwt.ErrFailedAuthentication
		} else if strings.Contains(err.Error(), "password not match") {
			return nil, jwt.ErrFailedAuthentication
		} else {
			return nil, jwt.ErrFailedAuthentication
		}
	}
	return dbResult, nil
}

// SignUp godoc
//	@Summary	邮件注册
//	@Schemes
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		user	body		UserSignUpRequest	true	"User"
//	@Success	200		object	 models.Resp  注册成功
//	@Router		/signup [post]
func (u UserControllersImpl) SignUp(ctx *gin.Context) {
	var user *UserSignUpRequest
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		models.FailWithMessage(err.Error(), ctx)
		return
	}
	dbResult, err := u.userS.SignUp(&models.UserInfo{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			models.FailWithMessage(err.Error(), ctx)
			return
		} else {
			models.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	models.OkWithData(dbResult, ctx)
}

var _ UserControllers = &UserControllersImpl{}
