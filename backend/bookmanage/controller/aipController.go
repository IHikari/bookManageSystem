package controller

import (
	"bookmanage/common"
	"bookmanage/model"
	"bookmanage/response"
	"bookmanage/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Register 接口用于注册新用户。
// @Summary 注册新用户
// @Description 从请求体中解析用户信息，并尝试注册用户。如果注册成功，返回成功消息；否则返回失败消息。
// @Accept json
// @Produce json
// @Param schoolcard body string true "学号"
// @Param name body string true "用户名"
// @Param password body string true "密码"
// @Router /reg [post]
func Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		fmt.Println(err)
		response.Fail(ctx, "数据传输失败", nil)
		return
	}
	flag := service.Register(&user)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}

	response.Success(ctx, nil, flag.Message)

}

// Login 接口用于用户登录。
// @Summary 用户登录
// @Description 从请求体中解析用户信息，并尝试进行用户登录。如果登录成功，返回用户的身份验证令牌；否则返回失败消息。
// @Accept json
// @Produce json
// @Param schoolcard body string true "学号"
// @Param password body string true "密码"
// @Router /login [post]
func Login(ctx *gin.Context) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		response.Fail(ctx, "数据传输失败", nil)
		return
	}
	flag := service.Login(&user)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	token, _ := common.ReleaseToken(user)
	response.Success(ctx, map[string]string{"token": token}, flag.Message)
}
