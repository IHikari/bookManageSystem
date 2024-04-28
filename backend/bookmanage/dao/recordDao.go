package dao

import (
	"bookmanage/common"
	"bookmanage/model"
	"bookmanage/utils"
	"gorm.io/gorm"
)

func CreateRecord(record *model.Record, db *gorm.DB) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	if err := common.DB.Create(record).Error; err != nil {
		db.Rollback()
		flag.Status = false
		flag.Message = "创建记录失败"
		return flag
	}
	return flag
}
