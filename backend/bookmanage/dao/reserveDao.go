package dao

import (
	"bookmanage/common"
	"bookmanage/model"
	"bookmanage/utils"
	"gorm.io/gorm"
	"time"
)

func ReserveBorrow(reserve *model.Reserve, db *gorm.DB) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	var temp []model.Reserve
	common.DB.Where("school_id=?", reserve.SchoolId).Find(&temp)
	for _, v := range temp {
		if v.BookId == reserve.BookId {
			v.BookNum += reserve.BookNum
			v.UpdatedAt = time.Now()
			err := common.DB.Select("book_num", "updated_at").Updates(v).Error
			if err != nil {
				db.Rollback()
				flag.Status = false
				flag.Message = "更新保留记录失败"
				return flag
			}
			flag.Message = "操作成功"
			return flag
		}
	}
	reserve.CreatedAt = time.Now()
	reserve.UpdatedAt = time.Now()
	if err := common.DB.Create(reserve).Error; err != nil {
		db.Rollback()
		flag.Status = false
		flag.Message = "新增保留记录失败"
		return flag
	}
	flag.Message = "借书成功"
	return flag
}

func ReserveLend(reserve *model.Reserve, db *gorm.DB) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	var temp []model.Reserve
	common.DB.Where("school_id=?", reserve.SchoolId).Find(&temp)
	for _, v := range temp {
		if v.BookId == reserve.BookId {
			v.BookNum -= reserve.BookNum
			if v.BookNum == 0 {
				if err := common.DB.Where("book_id=? and school_id=?", reserve.BookId, reserve.SchoolId).Delete(&model.Reserve{}); err != nil {
					db.Rollback()
					flag.Status = false
					flag.Message = "删除记录失败"
					return flag
				}
				continue
			}
			v.UpdatedAt = time.Now()
			err := common.DB.Select("book_num", "updated_at").Updates(v).Error
			if err != nil {
				db.Rollback()
				flag.Status = false
				flag.Message = "更新保留记录失败"
				return flag
			}
		}
	}
	flag.Message = "还书成功"
	return flag
}
