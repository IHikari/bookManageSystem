package test

import (
	"bookmanage/model"
	"bookmanage/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestServer(t *testing.T) {

	r := gin.Default()

	r.PUT("/user/info", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			fmt.Println(err)
			response.Fail(ctx, "数据传输失败", nil)
			return
		}
	})
	r.Run()
}
