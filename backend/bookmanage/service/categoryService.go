package service

import (
	"bookmanage/dao"
	"bookmanage/model"
	"bookmanage/utils"
)

func CategoryList(category *[]model.Category) utils.Flag {
	return dao.CategoryList(category)
}

func CategoryAdd(category *model.Category) utils.Flag {
	return dao.CategoryAdd(category)
}

func CategoryUpdate(category *model.Category) utils.Flag {
	return dao.CategoryUpdate(category)
}

func CategoryDelete(id int) utils.Flag {
	return dao.CategoryDelete(id)
}
