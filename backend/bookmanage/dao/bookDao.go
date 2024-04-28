package dao

import (
	"bookmanage/common"
	"bookmanage/dto"
	"bookmanage/model"
	"bookmanage/utils"
	"time"
)

func BookList(result *[]model.Book, pagination utils.Pagination, cateIdStr string, input string) (utils.Flag, []dto.BookCategory, int64) {
	var bookcategory []dto.BookCategory
	flag, total := utils.Paginate(common.DB, result, pagination, cateIdStr, input)
	if !flag.Status {
		return flag, bookcategory, 0
	}
	bookcategory = make([]dto.BookCategory, len(*result))
	var category model.Category
	for i := 0; i < len(*result); i++ {

		if err := common.DB.Where("id = ?", (*result)[i].CateId).First(&category).Error; err != nil {
			flag.Status = false
			flag.Message = "查询类别表失败"
			return flag, bookcategory, 0
		}
		(bookcategory)[i].ID = (*result)[i].ID
		(bookcategory)[i].BookName = (*result)[i].BookName
		(bookcategory)[i].Publisher = (*result)[i].Publisher
		(bookcategory)[i].Author = (*result)[i].Author
		(bookcategory)[i].CateId = (*result)[i].CateId
		(bookcategory)[i].Image = (*result)[i].Image
		(bookcategory)[i].Number = (*result)[i].Number
		(bookcategory)[i].CateName = category.Name

		category = model.Category{}
	}
	return flag, bookcategory, total
}
func GetByName(bookName string) model.Book {
	var result model.Book
	common.DB.Where("book_name=?", bookName).First(&result)
	return result
}

func BookAdd(book *model.Book) utils.Flag {
	// 标志
	var flag utils.Flag
	flag.Status = true
	// 判断该书是否存在
	var temp model.Book
	booknameFlag := common.DB.Where("book_name =?", book.BookName).Find(&temp).RowsAffected
	authorFlag := common.DB.Where("author =?", book.Author).Find(&temp).RowsAffected
	publisherFlag := common.DB.Where("publisher =?", book.Publisher).Find(&temp).RowsAffected
	cateIdFlag := common.DB.Where("cate_id =?", book.CateId).Find(&temp).RowsAffected

	if booknameFlag == 1 && authorFlag == 1 && publisherFlag == 1 && cateIdFlag == 1 {
		flag.Status = false
		flag.Message = "该图书已存在"
		return flag
	}
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	if err := common.DB.Create(book).Error; err != nil {
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	flag.Message = "创建成功"
	return flag
}

func BookGetDetail(result *model.Book) utils.Flag {
	// 标志
	var flag utils.Flag
	flag.Status = true
	if err := common.DB.Where("id=?", result.ID).First(result).Error; err != nil {
		flag.Status = false
		flag.Message = "查询数据库失败"
		return flag
	}
	flag.Message = "操作成功"
	return flag

}
func BookEdit(book *model.Book) (utils.Flag, model.Book) {
	// 标志
	var flag utils.Flag
	flag.Status = true
	book.UpdatedAt = time.Now()
	if err := common.DB.Where("id=?", book.ID).Select("*").Omit("created_at").Updates(*book).Error; err != nil {
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag, *book
	}
	flag.Message = "编辑成功"
	return flag, *book
}

func BookDelete(id uint) utils.Flag {
	// 标志
	var flag utils.Flag
	flag.Status = true
	tx := common.DB.Begin()
	err := common.DB.Where("book_id=?", id).Delete(&model.Record{}).Error
	if err != nil {
		tx.Rollback()
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	err = common.DB.Where("book_id=?", id).Delete(&model.Reserve{}).Error
	if err != nil {
		tx.Rollback()
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	if err := common.DB.Where("id=?", id).Delete(&model.Book{}).Error; err != nil {
		tx.Rollback()
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	flag.Message = "删除成功"
	return flag
}

func BookUserList(result *[]dto.BookCategory) {
	for i := 0; i < len(*result); i++ {
		var temp model.Book
		common.DB.Where("id=?", (*result)[i].ID).Find(&temp)
		(*result)[i].BookName = temp.BookName
		(*result)[i].Author = temp.Author
		(*result)[i].CateId = temp.CateId
		(*result)[i].Publisher = temp.Publisher
		(*result)[i].Image = temp.Image
		temp = model.Book{}
	}
}

// 书被借出
func BookBeBorrow(id, number uint) {
	var book model.Book
	book.ID = id
	common.DB.Find(&book)
	book.Number -= number
	book.UpdatedAt = time.Now()
	common.DB.Select("number", "updated_at").Updates(book)
}

// 书被归还
func BookBeLent(id, number uint) {
	var book model.Book
	book.ID = id
	common.DB.Find(&book)
	book.Number += number
	book.UpdatedAt = time.Now()
	common.DB.Select("number", "updated_at").Updates(book)
}
func UserBookGetDetail(result *model.Book) utils.Flag {
	// 标志
	var flag utils.Flag
	flag.Status = true
	if err := common.DB.Where("id=?", result.ID).First(result).Error; err != nil {
		flag.Status = false
		flag.Message = "查询数据库失败"
		return flag
	}
	flag.Message = "操作成功"
	return flag

}

func BookGetByIdToList(result *[]dto.RecordDto) {
	for i := 0; i < len(*result); i++ {
		var book model.Book
		common.DB.Where("id=?", (*result)[i].BookId).Find(&book)
		(*result)[i].BookName = book.BookName
		book = model.Book{}
	}
}
