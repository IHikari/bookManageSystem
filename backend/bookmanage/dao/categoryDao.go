package dao

import (
	"bookmanage/common"
	"bookmanage/dto"
	"bookmanage/model"
	"bookmanage/utils"
	"time"
)

func CategoryList(category *[]model.Category) utils.Flag {
	var flag utils.Flag
	flag.Status = true

	if err := common.DB.Find(category).Error; err != nil {
		flag.Status = false
		flag.Message = "查询数据库失败"
		return flag
	}
	flag.Message = "获取图书分类列表成功"
	return flag
}

func CategoryAdd(category *model.Category) utils.Flag {
	var flag utils.Flag
	var temp model.Category
	flag.Status = true

	if err := common.DB.Find(&temp).Error; err != nil {
		flag.Status = false
		flag.Message = "查询数据库失败"
		return flag
	}
	if temp.ID == category.ID {
		flag.Status = false
		flag.Message = "改类别编号已存在"
		return flag
	}
	if temp.Name == category.Name {
		flag.Status = false
		flag.Message = "改类别名称已存在"
		return flag
	}
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	common.DB.Create(category)
	flag.Message = "添加成功"
	return flag
}
func CategoryUpdate(category *model.Category) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	category.UpdatedAt = time.Now()
	err := common.DB.Model(&model.Category{}).Where("id=?", category.ID).Updates(category).Error
	if err != nil {
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	flag.Message = "修改成功"
	return flag
}
func CategoryDelete(id int) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	tx := common.DB.Begin()
	if err := common.DB.Where("cate_id=?", id).Delete(&model.Book{}).Error; err != nil {
		tx.Rollback()
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	if err := common.DB.Where("id = ?", id).Delete(&model.Category{}).Error; err != nil {
		tx.Rollback()
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	flag.Message = "删除成功"
	return flag
}

func CategoryUserList(result *[]dto.BookCategory) {
	for i := 0; i < len(*result); i++ {
		var temp model.Category
		common.DB.Where("id=?", (*result)[i].CateId).Find(&temp)
		(*result)[i].CateName = temp.Name
		temp = model.Category{}
	}
}
