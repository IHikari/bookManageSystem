package service

import (
	"bookmanage/dao"
	"bookmanage/dto"
	"bookmanage/model"
	"bookmanage/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"path/filepath"
	"strconv"
)

func BookList(result *[]model.Book, pagination utils.Pagination, cateIdStr string, input string) (utils.Flag, []dto.BookCategory, int64) {
	return dao.BookList(result, pagination, cateIdStr, input)
}

func BookGetDetail(result *model.Book) utils.Flag {

	return dao.BookGetDetail(result)
}

// 添加图书
func BookAdd(bookName, author, cateIdStr, publisher, numberStr, remarks string, file *multipart.FileHeader, ctx *gin.Context) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	// 生成uuid
	u := uuid.New()
	// 文件后缀
	ext := filepath.Ext(file.Filename)
	// 新文件名为 UUID + 后缀
	newFileName := u.String() + ext
	path := fmt.Sprintf("D:/temp/%s", newFileName)
	err := ctx.SaveUploadedFile(file, path)
	if err != nil {
		flag.Status = false
		flag.Message = "文件写入失败"
		return flag
	}
	utils.UpFile(newFileName, path)

	// 转型
	cateId, _ := strconv.Atoi(cateIdStr)
	number, _ := strconv.Atoi(numberStr)

	// 实例化
	var book model.Book
	book.BookName = bookName
	book.Author = author
	book.CateId = uint(cateId)
	book.Publisher = publisher
	book.Image = "https://book-manage-system.oss-cn-hangzhou.aliyuncs.com/" + newFileName
	book.Number = uint(number)
	book.Remarks = remarks
	return dao.BookAdd(&book)
}

func BookEdit(id, bookName, author, cateIdStr, publisher, numberStr, remarks string, image any, ctx *gin.Context) (utils.Flag, model.Book) {

	var flag utils.Flag
	flag.Status = true
	// 实例化
	var book model.Book
	var url string
	if image.(string) == "" {
		// 获得表单数据
		file, err := ctx.FormFile("image")
		if err != nil {
			flag.Status = false
			flag.Message = "数据解析失败"
			return flag, book
		}
		// 生成uuid
		u := uuid.New()
		// 文件后缀
		ext := filepath.Ext(file.Filename)
		// 新文件名为 UUID + 后缀
		newFileName := u.String() + ext
		path := fmt.Sprintf("D:/temp/%s", newFileName)
		err = ctx.SaveUploadedFile(file, path)
		if err != nil {
			flag.Status = false
			flag.Message = "文件操作失败"
			return flag, book
		}
		utils.UpFile(newFileName, path)
		url = "https://book-manage-system.oss-cn-hangzhou.aliyuncs.com/" + newFileName
	}
	// 转型
	ID, _ := strconv.Atoi(id)
	cateId, _ := strconv.Atoi(cateIdStr)
	number, _ := strconv.Atoi(numberStr)

	book.ID = uint(ID)
	book.BookName = bookName
	book.Author = author
	book.CateId = uint(cateId)
	book.Publisher = publisher
	book.Image = url
	book.Number = uint(number)
	book.Remarks = remarks
	return dao.BookEdit(&book)
}

func BookDelete(id uint) utils.Flag {
	return dao.BookDelete(id)
}
