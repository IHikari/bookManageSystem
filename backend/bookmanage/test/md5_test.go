package test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	data := "123456"

	// 对数据进行MD5加密
	encrypted := EncryptMD5(data)
	fmt.Println("MD5加密后的值:", encrypted)

}
func EncryptMD5(data string) string {
	// 创建MD5哈希对象
	hash := md5.New()

	// 将数据写入哈希对象
	hash.Write([]byte(data))

	// 计算哈希值
	hashed := hash.Sum(nil)

	// 将哈希值转换为16进制字符串
	encrypted := hex.EncodeToString(hashed)

	return encrypted
}
