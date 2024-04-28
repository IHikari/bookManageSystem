package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Tomd(data string) string {
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
