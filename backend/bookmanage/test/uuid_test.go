package test

import (
	"fmt"
	"github.com/google/uuid"
	"path/filepath"
	"testing"
)

func TestUuid(t *testing.T) {
	// 原始文件名
	fileName := "example.txt"

	// 获取原始文件名的后缀
	ext := filepath.Ext(fileName)

	// 生成 UUID
	u := uuid.New()

	// 新文件名为 UUID + 后缀
	newFileName := u.String() + ext

	// 打印新文件名
	fmt.Println("New file name:", newFileName)
}
