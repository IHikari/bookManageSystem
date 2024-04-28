package controller

import (
	"bookmanage/model"
	"bookmanage/response"
	"bookmanage/service"
	"bookmanage/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BookList(ctx *gin.Context) {

	var result []model.Book

	pageStr := ctx.DefaultQuery("pagenum", "1")
	pageSizeStr := ctx.DefaultQuery("pagesize", "5")
	cateIdStr := ctx.Query("cate_id")
	// 查询书名输入
	input := ctx.Query("input")

	// 将字符串转换为整数
	pagenum, _ := strconv.Atoi(pageStr)
	pagesize, _ := strconv.Atoi(pageSizeStr)
	pagination := utils.Pagination{Page: pagenum, PageSize: pagesize}
	flag, bookCategory, total := service.BookList(&result, pagination, cateIdStr, input)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取图书列表成功", "data": bookCategory, "total": total})
}

// 添加图书
func BookAdd(ctx *gin.Context) {
	// 获得表单数据
	file, err := ctx.FormFile("image")
	if err != nil {
		response.Fail(ctx, "数据解析失败", nil)
		return
	}
	bookName := ctx.PostForm("bookname")
	author := ctx.PostForm("author")
	cateIdStr := ctx.PostForm("cate_id")
	publisher := ctx.PostForm("publisher")
	numberStr := ctx.PostForm("number")
	remarks := ctx.PostForm("remarks")
	flag := service.BookAdd(bookName, author, cateIdStr, publisher, numberStr, remarks, file, ctx)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
	}
	response.Success(ctx, nil, flag.Message)
}

// 根据id查询图书
func BookGetDetail(ctx *gin.Context) {

	Id := ctx.Query("id")
	id, _ := strconv.Atoi(Id)
	var book model.Book
	book.ID = uint(id)

	flag := service.BookGetDetail(&book)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}

	response.Success(ctx, book, flag.Message)
}

// 编辑图书
func BookEdit(ctx *gin.Context) {

	id := ctx.PostForm("ID")
	bookName := ctx.PostForm("bookname")
	author := ctx.PostForm("author")
	cateIdStr := ctx.PostForm("cate_id")
	publisher := ctx.PostForm("publisher")
	numberStr := ctx.PostForm("number")
	remarks := ctx.PostForm("remarks")
	image := ctx.PostForm("image")

	flag, book := service.BookEdit(id, bookName, author, cateIdStr, publisher, numberStr, remarks, image, ctx)
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}

	response.Success(ctx, book, flag.Message)
}
func BookDelete(ctx *gin.Context) {

	id := ctx.Query("id")
	ID, _ := strconv.Atoi(id)
	flag := service.BookDelete(uint(ID))
	if !flag.Status {
		response.Fail(ctx, flag.Message, nil)
		return
	}
	response.Success(ctx, nil, flag.Message)
}
