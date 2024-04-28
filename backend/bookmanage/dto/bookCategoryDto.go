package dto

import "bookmanage/model"

type BookCategory struct {
	model.Book
	CateName string `json:"cate_name"`
}
