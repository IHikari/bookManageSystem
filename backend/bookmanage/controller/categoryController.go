package controller

import (
	"bookmanage/model"
	"bookmanage/response"
	"bookmanage/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CategoryList 加载类别列表
func CategoryList(ctx *gin.Context) {
	var category []model.Category

	flag := service.CategoryList(&category)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, category, flag.Message)
}

// CategoryAdd 添加类别
func CategoryAdd(ctx *gin.Context) {
	var category model.Category
	if err := ctx.BindJSON(&category); err != nil {
		response.Fail(ctx, "数据传输失败", nil)
		return
	}
	flag := service.CategoryAdd(&category)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, nil, flag.Message)
}

// CategoryUpdate 更新类别
func CategoryUpdate(ctx *gin.Context) {
	var category model.Category
	if err := ctx.BindJSON(&category); err != nil {
		response.Fail(ctx, "数据传输失败", nil)
		return
	}
	flag := service.CategoryUpdate(&category)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, nil, flag.Message)
}

// CategoryDelete 删除类别
func CategoryDelete(ctx *gin.Context) {
	ID := ctx.Query("id")
	id, _ := strconv.Atoi(ID)
	flag := service.CategoryDelete(id)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, nil, flag.Message)
}
