package test

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	// 定义一个测试字符串
	str := 1

	s := fmt.Sprintf("%d%%", str)
	fmt.Println(s)
}
