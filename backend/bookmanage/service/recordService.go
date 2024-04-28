package service

import (
	"bookmanage/common"
	"bookmanage/dao"
	"bookmanage/dto"
	"bookmanage/model"
	"bookmanage/utils"
	"strconv"
	"time"
)

func BorrowBook(schoolIdStr, bookIdStr, bookNumStr string) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	// 转为整形
	school_id, _ := strconv.Atoi(schoolIdStr)
	book_id, _ := strconv.Atoi(bookIdStr)
	book_num, _ := strconv.Atoi(bookNumStr)

	tx := common.DB.Begin()
	dao.BookBeBorrow(uint(book_id), uint(book_num))

	var record model.Record
	record.SchoolId = uint(school_id)
	record.BookId = uint(book_id)
	record.BookNum = uint(book_num)
	record.IsBorrow = true
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()

	flag = dao.CreateRecord(&record, tx)
	if !flag.Status {
		return flag
	}
	var reserve model.Reserve
	reserve.SchoolId = record.SchoolId
	reserve.BookId = record.BookId
	reserve.BookNum = record.BookNum
	return dao.ReserveBorrow(&reserve, tx)
}

func LendBook(schoolIdStr, bookIdStr, bookNumStr string) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	// 转为整形
	school_id, _ := strconv.Atoi(schoolIdStr)
	book_id, _ := strconv.Atoi(bookIdStr)
	book_num, _ := strconv.Atoi(bookNumStr)

	tx := common.DB.Begin()
	dao.BookBeLent(uint(book_id), uint(book_num))

	var record model.Record
	record.SchoolId = uint(school_id)
	record.BookId = uint(book_id)
	record.BookNum = uint(book_num)
	record.IsBorrow = false
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()

	flag = dao.CreateRecord(&record, tx)
	if !flag.Status {
		return flag
	}
	var reserve model.Reserve
	reserve.SchoolId = record.SchoolId
	reserve.BookId = record.BookId
	reserve.BookNum = record.BookNum
	return dao.ReserveLend(&reserve, tx)
}

func RecordList(result *[]dto.RecordDto, school_id uint) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	var user model.User
	var record []model.Record
	common.DB.Where("id=?", school_id).Find(&user)
	if user.Permission == true {
		common.DB.Find(&record)
	} else {
		common.DB.Where("school_id=?", school_id).Find(&record)
	}
	if len(record) == 0 {
		flag.Status = true
		flag.Message = "无借阅记录"
		return flag
	}
	*result = make([]dto.RecordDto, len(record))
	for i := 0; i < len(record); i++ {
		(*result)[i].SchoolId = record[i].SchoolId
		(*result)[i].BookId = record[i].BookId
		(*result)[i].BookNum = record[i].BookNum
		(*result)[i].IsBorrow = record[i].IsBorrow
		(*result)[i].CreatedAt = record[i].CreatedAt
	}
	// 查user表
	dao.UserGetByIdToList(result)
	// 查图书表
	dao.BookGetByIdToList(result)
	flag.Message = "查阅成功"
	return flag
}
