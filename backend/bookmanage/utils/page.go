package utils

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type Pagination struct {
	Page     int //当前页码
	PageSize int //每页大小
}

// book分页
func Paginate(query *gorm.DB, result interface{}, pagination Pagination, cateIdStr string, input string) (Flag, int64) {
	var flag Flag
	flag.Status = true
	// 计算偏移量
	offset := (pagination.Page - 1) * pagination.PageSize
	s := fmt.Sprintf("%%%s%%", input)
	// 总记录数
	var total int64
	if cateIdStr == "" {
		total = query.Where("book_name like ?", s).Find(result).RowsAffected
		// 查询数据库
		err := query.Where("book_name like ?", s).Offset(offset).Limit(pagination.PageSize).Find(result).Error
		if err != nil {
			flag.Status = false
			flag.Message = "查询数据库失败"
			return flag, 0
		}
	} else {
		cateId, _ := strconv.Atoi(cateIdStr)
		total = query.Where("book_name like ? and cate_id = ?", s, cateId).Find(result).RowsAffected
		// 查询数据库
		err := query.Where("book_name like ? and cate_id = ?", s, cateId).Offset(offset).Limit(pagination.PageSize).Where("cate_id = ?", cateId).Find(result).Error
		if err != nil {
			flag.Status = false
			flag.Message = "查询数据库失败"
			return flag, 0
		}
	}
	flag.Message = "操作成功"
	return flag, total
}
