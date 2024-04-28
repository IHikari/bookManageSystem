package dao

import (
	"bookmanage/common"
	"bookmanage/model"
	"bookmanage/utils"
)

func Register(user *model.User) utils.Flag {
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
	flag.Message = "注册成功！"
	return flag
}

func Login(user *model.User) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	var temp model.User
	if common.DB.Where("id=?", user.ID).First(&temp).RowsAffected == 0 {
		flag.Status = false
		flag.Message = "该用户不存在,请先注册！"
		return flag
	}
	if user.Password != temp.Password {
		flag.Status = false
		flag.Message = "密码错误，请重新输入！"
		return flag
	}
	flag.Message = "登录成功！"
	return flag

}
