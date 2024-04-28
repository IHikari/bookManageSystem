package service

import (
	"bookmanage/dao"
	"bookmanage/model"
	"bookmanage/utils"
	"strconv"
	"time"
)

func UserList(result *[]model.User, pageStr, pageSizeStr, input string) (utils.Flag, int64) {

	// 将字符串转换为整数
	pagenum, _ := strconv.Atoi(pageStr)
	pagesize, _ := strconv.Atoi(pageSizeStr)
	pagination := utils.Pagination{Page: pagenum, PageSize: pagesize}

	_, err := strconv.Atoi(input)
	if err == nil {
		// input是学工号,或空
		id, _ := strconv.Atoi(input)
		return dao.UserListById(result, pagination, uint(id))
	} else {
		// input是姓名
		name := input
		return dao.UserListByname(result, pagination, name)
	}

}

func UserEditRoot(user *model.User) utils.Flag {
	user.Password = utils.Tomd(user.Password)
	user.UpdatedAt = time.Now()
	return dao.UserEditRoot(user)
}

func UserAdd(user *model.User) utils.Flag {
	user.Password = utils.Tomd(user.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return dao.UserAdd(user)
}

func UserDelete(id uint) utils.Flag {
	return dao.UserDelete(id)
}
