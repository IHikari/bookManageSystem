package controller

import (
	"bookmanage/model"
	"bookmanage/response"
	"bookmanage/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserList(ctx *gin.Context) {
	var result []model.User

	pageStr := ctx.DefaultQuery("pagenum", "1")
	pageSizeStr := ctx.DefaultQuery("pagesize", "5")
	// 查询学工号或姓名
	input := ctx.Query("input")

	flag, total := service.UserList(&result, pageStr, pageSizeStr, input)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取用户列表成功", "data": result, "total": total})
}

// 编辑用户密码和权限
func UserEditRoot(ctx *gin.Context) {

	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		fmt.Println(err)
		response.Fail(ctx, "数据传输失败", nil)
		return
	}
	flag := service.UserEditRoot(&user)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, user, flag.Message)

}

// 新增用户
func UserAdd(ctx *gin.Context) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		fmt.Println(err)
		response.Fail(ctx, "数据传输失败", nil)
		return
	}
	flag := service.UserAdd(&user)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}

	response.Success(ctx, nil, flag.Message)

}

// //
func UserDelete(ctx *gin.Context) {
	Id := ctx.Query("id")
	id, _ := strconv.Atoi(Id)
	flag := service.UserDelete(uint(id))
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, nil, flag.Message)
}
