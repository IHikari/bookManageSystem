package model

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	SchoolId uint `json:"school_id"` // 借还书的人的学工号
	BookId   uint `json:"book_id"`   // 借还的书的编号
	BookNum  uint `json:"book_num"`  // 借还书的数量
	IsBorrow bool `json:"is_borrow"` // 借书：true 还书：false
}
