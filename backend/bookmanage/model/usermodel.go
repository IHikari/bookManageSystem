package model

import "gorm.io/gorm"

type User struct {
	Name       string `gorm:"not null" json:"name"`
	Nickname   string `gorm:"type:varchar(50)" json:"nickname"`
	Password   string `gorm:"not null" json:"password"`
	Permission bool   `gorm:"default:0" json:"permission""`      // 0- 普通用户  1- 管理员
	UserPic    string `gorm:"type:varchar(255)" json:"user_pic"` //头像
	Email      string `gorm:"type:varchar(50)" json:"email"`
	gorm.Model
}
