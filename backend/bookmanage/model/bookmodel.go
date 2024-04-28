package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookName  string `json:"bookname" gorm:"type:varchar(50);not null"`
	Image     string `json:"image" gorm:"type:varchar(255)"` // 图片
	Author    string `json:"author" gorm:"type:varchar(50);not null"`
	CateId    uint   `json:"cate_id" gorm:"default:0"`          //类别ID
	Publisher string `json:"publisher" gorm:"type:varchar(50)"` //出版社
	Number    uint   `json:"number" gorm:"default:1"`           //数量
	Remarks   string `json:"remarks" gorm:"type:varchar(255)"`  //描述
}
