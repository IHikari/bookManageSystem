package controller

import (
	"bookmanage/dto"
	"bookmanage/response"
	"bookmanage/service"
	"github.com/gin-gonic/gin"
)

func BorrowBook(ctx *gin.Context) {
	schoolIdStr := ctx.PostForm("school_id")
	bookIdStr := ctx.PostForm("book_id")
	bookNumStr := ctx.PostForm("book_num")

	flag := service.BorrowBook(schoolIdStr, bookIdStr, bookNumStr)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, nil, flag.Message)
}

func LendBook(ctx *gin.Context) {
	schoolIdStr := ctx.PostForm("school_id")
	bookIdStr := ctx.PostForm("book_id")
	bookNumStr := ctx.PostForm("book_num")

	flag := service.LendBook(schoolIdStr, bookIdStr, bookNumStr)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, nil, flag.Message)
}

func RecordList(ctx *gin.Context) {
	schoolId, err := ctx.Get("ID")
	if !err {
		response.Fail(ctx, "令牌不存在", nil)
		return
	}
	var result []dto.RecordDto

	flag := service.RecordList(&result, schoolId.(uint))
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, result, flag.Message)
}
