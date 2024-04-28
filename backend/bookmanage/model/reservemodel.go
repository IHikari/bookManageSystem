package model

import "gorm.io/gorm"

type Reserve struct {
	gorm.Model
	SchoolId uint `json:"school_id"` // 人
	BookId   uint `json:"book_id"`   // 数
	BookNum  uint `json:"book_num"`  // 保留的本书
}
