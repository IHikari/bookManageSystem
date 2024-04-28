package dao

import (
	"bookmanage/common"
	"bookmanage/dto"
	"bookmanage/model"
	"bookmanage/utils"
	"fmt"
)

func UserListById(result *[]model.User, pagination utils.Pagination, id uint) (utils.Flag, int64) {
	var flag utils.Flag
	flag.Status = true
	// 计算偏移量
	offset := (pagination.Page - 1) * pagination.PageSize
	// 根据主键模糊匹配
	s := fmt.Sprintf("%d%%", id)
	var total int64
	total = common.DB.Where("id like ?", s).Find(result).RowsAffected
	// 查询数据库
	err := common.DB.Where("id like ?", s).Offset(offset).Limit(pagination.PageSize).Find(result).Error
	if err != nil {
		flag.Status = false
		flag.Message = "查询数据库失败"
		return flag, 0
	}
	flag.Message = "操作成功"
	return flag, total
}

func UserListByname(result *[]model.User, pagination utils.Pagination, name string) (utils.Flag, int64) {
	var flag utils.Flag
	flag.Status = true
	// 计算偏移量
	offset := (pagination.Page - 1) * pagination.PageSize
	// 名字模糊匹配
	s := fmt.Sprintf("%%%s%%", name)
	// 总记录数
	var total int64

	total = common.DB.Where("name like ?", s).Find(result).RowsAffected
	// 查询数据库
	err := common.DB.Where("name like ?", s).Offset(offset).Limit(pagination.PageSize).Find(result).Error
	if err != nil {
		flag.Status = false
		flag.Message = "查询数据库失败"
		return flag, 0
	}

	flag.Message = "操作成功"
	return flag, total
}

func UserEditRoot(user *model.User) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	err := common.DB.Where("id=?", user.ID).Select("password", "permission", "updated_at").Updates(user).Error
	if err != nil {
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	flag.Message = "修改成功"
	return flag

}

func UserAdd(user *model.User) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	if common.DB.Where("id=?", user.ID).First(user).RowsAffected == 1 {
		flag.Status = false
		flag.Message = "该学工号已被注册！"
		return flag
	}
	err := common.DB.Create(user).Error
	if err != nil {
		flag.Status = false
		flag.Message = "创建数据失败！"
		return flag
	}
	flag.Message = "新增成功！"
	return flag
}

func UserGetByIdToList(result *[]dto.RecordDto) {
	for i := 0; i < len(*result); i++ {
		var user model.User
		common.DB.Where("id=?", (*result)[i].SchoolId).Find(&user)
		(*result)[i].Name = user.Name
		user = model.User{}
	}
}

func UserDelete(id uint) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	tx := common.DB.Begin()
	err := common.DB.Where("school_id=?", id).Delete(&model.Record{}).Error
	if err != nil {
		tx.Rollback()
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	err = common.DB.Where("school_id=?", id).Delete(&model.Reserve{}).Error
	if err != nil {
		tx.Rollback()
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	err = common.DB.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		tx.Rollback()
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	flag.Message = "删除成功"
	return flag
}
