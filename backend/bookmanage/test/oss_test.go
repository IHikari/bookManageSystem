package test

import (
	"bookmanage/utils"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"testing"
)

func TestOss(t *testing.T) {
	fmt.Println("OSS Go SDK Version: ", oss.Version)

	utils.UpFile("1.jpg", "D:\\project\\Go_WorkSpace\\src\\bookmanage\\img\\6da1e84e-90a7-4610-acfe-2721d169b896.jpg")

}
