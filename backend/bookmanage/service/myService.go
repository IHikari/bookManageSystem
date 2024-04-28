package service

import (
	"bookmanage/common"
	"bookmanage/dao"
	"bookmanage/dto"
	"bookmanage/model"
	"bookmanage/utils"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
	"time"
)

func UserInfo(user *model.User) utils.Flag {
	return dao.UserInfo(user)
}

func UserEdit(user *model.User) utils.Flag {
	user.UpdatedAt = time.Now()
	return dao.UserEdit(user)
}

func UserEditAvatar(user *model.User, data []byte) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	base64String := string(data)
	// 将前缀和文件内容分离
	prefixStr := strings.Split(base64String, ";base64,")[0]
	encodedData := strings.Split(base64String, ";base64,")[1]
	// 从前缀中获得文件类型
	fileType := strings.Split(prefixStr, "/")[1]
	// 将文件内容解码为字节数组
	decoded, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		flag.Status = false
		flag.Message = "文件解码失败"
		return flag
	}
	// 生成文件名
	u := uuid.New()
	// 生成文件路径
	fileName := fmt.Sprintf("%s.%s", u.String(), fileType)
	path := fmt.Sprintf("D:/temp/%s", fileName)
	// 将解码后的字节数组写入到本地文件
	err = os.WriteFile(path, decoded, 0644)
	if err != nil {
		flag.Status = false
		flag.Message = "写入文件时出错"
		return flag
	}
	// 将文件上传至阿里云oss
	utils.UpFile(fileName, path)
	// 获得用户ID
	user.UserPic = "https://book-manage-system.oss-cn-hangzhou.aliyuncs.com/" + fileName
	user.UpdatedAt = time.Now()
	return dao.UserEditAvatar(user)
}

func UserUpdatePwd(user *model.User, rePassword *dto.RePassword) utils.Flag {
	var flag utils.Flag
	flag.Status = true
	if rePassword.NewPwd != rePassword.RePwd {
		flag.Status = false
		flag.Message = "两次密码不一致"
		return flag
	}

	rePassword.OldPwd = utils.Tomd(rePassword.OldPwd)
	rePassword.NewPwd = utils.Tomd(rePassword.NewPwd)
	rePassword.RePwd = utils.Tomd(rePassword.RePwd)
	user.UpdatedAt = time.Now()
	return dao.UserUpdatePwd(user, rePassword)
}

func UserBookList(result *[]dto.BookCategory, user *model.User) utils.Flag {
	var flag utils.Flag
	flag.Status = true

	var reserve []model.Reserve
	common.DB.Where("school_id=?", user.ID).Find(&reserve)

	if len(reserve) == 0 {
		flag.Message = "没有数据"
		return flag
	}

	*result = make([]dto.BookCategory, len(reserve))
	for i := 0; i < len(reserve); i++ {
		(*result)[i].ID = reserve[i].BookId
		(*result)[i].Number = reserve[i].BookNum
	}
	dao.BookUserList(result)
	dao.CategoryUserList(result)
	flag.Message = "获取成功"
	return flag

}

func UserBookGetDetail(result *model.Book, reserve *model.Reserve) utils.Flag {

	common.DB.Where("school_id=? and book_id=?", reserve.SchoolId, reserve.BookId).Find(reserve)

	flag := dao.UserBookGetDetail(result)
	if flag.Status == false {
		return flag
	}
	result.Number = reserve.BookNum
	return flag
}
