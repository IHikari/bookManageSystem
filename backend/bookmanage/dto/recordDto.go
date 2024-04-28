package dto

import "time"

type RecordDto struct {
	SchoolId  uint      `json:"school_id"`
	Name      string    `json:"name"`
	BookId    uint      `json:"book_id"`
	BookName  string    `json:"book_name"`
	CreatedAt time.Time `json:"created_at"`
	BookNum   uint      `json:"book_num"`
	IsBorrow  bool      `json:"is_borrow"`
}
