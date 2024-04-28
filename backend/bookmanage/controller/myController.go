package controller

import (
	"bookmanage/dto"
	"bookmanage/model"
	"bookmanage/response"
	"bookmanage/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"strings"
)

// 返回用户详情信息
func UserInfo(ctx *gin.Context) {

	ID, Flag := ctx.Get("ID")
	if !Flag {
		response.Fail(ctx, "解析令牌出错", nil)
		return
	}
	var user model.User
	user.ID = ID.(uint)
	flag := service.UserInfo(&user)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, user, flag.Message)

}

// 编辑用户基本信息
func UserEdit(ctx *gin.Context) {

	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		fmt.Println(err)
		response.Fail(ctx, "数据传输失败", nil)
		return
	}
	flag := service.UserEdit(&user)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, user, flag.Message)

}

// 编辑用户头像
func UserEditAvatar(ctx *gin.Context) {

	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		response.Fail(ctx, "请求体阅读失败", nil)
		return
	}
	if strings.Split(string(data), ":")[0] == "https" {
		response.Fail(ctx, "请更换图片", nil)
		return
	}
	ID, Flag := ctx.Get("ID")
	if !Flag {
		response.Fail(ctx, "解析令牌出错", nil)
		return
	}
	var user model.User
	user.ID = ID.(uint)
	flag := service.UserEditAvatar(&user, data)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, user, flag.Message)

}

func UserUpdatePwd(ctx *gin.Context) {

	var rePassword dto.RePassword
	if err := ctx.BindJSON(&rePassword); err != nil {
		fmt.Println(err)
		response.Fail(ctx, "数据传输失败", nil)
		return
	}
	ID, Flag := ctx.Get("ID")
	if !Flag {
		response.Fail(ctx, "解析令牌出错", nil)
		return
	}
	var user model.User
	user.ID = ID.(uint)
	flag := service.UserUpdatePwd(&user, &rePassword)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, user, flag.Message)

}

func UserBookList(ctx *gin.Context) {

	Id, err := ctx.Get("ID")
	if !err {
		response.Fail(ctx, "令牌不存在", nil)
		return
	}
	var result []dto.BookCategory
	var user model.User
	user.ID = Id.(uint)
	flag := service.UserBookList(&result, &user)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, result, flag.Message)
}

func UserBook(ctx *gin.Context) {
	schoolId, err := ctx.Get("ID")
	if !err {
		response.Fail(ctx, "令牌不存在", nil)
		return
	}

	Book_Id := ctx.Query("id")
	id, _ := strconv.Atoi(Book_Id)
	var book model.Book
	book.ID = uint(id)
	var reserve model.Reserve
	reserve.BookId = book.ID
	reserve.SchoolId = schoolId.(uint)
	flag := service.UserBookGetDetail(&book, &reserve)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}

	response.Success(ctx, book, flag.Message)
}
