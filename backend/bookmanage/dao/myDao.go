package dao

import (
	"bookmanage/common"
	"bookmanage/dto"
	"bookmanage/model"
	"bookmanage/utils"
	"fmt"
)

func UserInfo(user *model.User) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	fmt.Println(user)
	err := common.DB.Find(&user).Error
	if err != nil {
		flag.Status = false
		flag.Message = "查询数据库失败"
		return flag
	}
	fmt.Println(user)
	flag.Message = "请求成功"
	return flag

}
func UserEdit(user *model.User) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	err := common.DB.Where("id=?", user.ID).Select("nickname", "email", "updated_at").Updates(user).Error
	if err != nil {
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	flag.Message = "编辑成功"
	return flag

}
func UserEditAvatar(user *model.User) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	err := common.DB.Where("id=?", user.ID).Select("user_pic", "updated_at").Updates(user).Error
	if err != nil {
		flag.Status = false
		flag.Message = "操作数据库失败"
		return flag
	}
	flag.Message = "编辑成功"
	return flag

}

func UserUpdatePwd(user *model.User, rePassword *dto.RePassword) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	if err := common.DB.Find(user).Error; err != nil {
		flag.Status = false
		flag.Message = "查询数据库失败"
		return flag
	}
	if user.Password != rePassword.OldPwd {
		flag.Status = false
		flag.Message = "原密码错误"
		return flag
	}
	user.Password = rePassword.NewPwd
	common.DB.Select("password", "created_at").Updates(user)
	flag.Message = "重置密码成功"
	return flag
}
