package service

import (
	"bookmanage/dao"
	"bookmanage/model"
	"bookmanage/utils"
	"time"
)

func Register(user *model.User) utils.Flag {
	user.Password = utils.Tomd(user.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return dao.Register(user)
}

func Login(user *model.User) utils.Flag {
	user.Password = utils.Tomd(user.Password)
	return dao.Login(user)
}
