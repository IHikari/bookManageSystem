package router

import (
	"bookmanage/controller"
	"bookmanage/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {

	r := gin.Default()

	// 登录注册
	r.POST("/reg", controller.Register)
	r.POST("/login", controller.Login)

	r.Use(middleware.AuthMiddleware())
	// 个人信息
	my := r.Group("my")
	my.GET("/userinfo", controller.UserInfo)
	my.PUT("/userinfo", controller.UserEdit)
	my.PATCH("/avatar", controller.UserEditAvatar)
	my.PATCH("/updatepwd", controller.UserUpdatePwd)
	my.GET("/booklist", controller.UserBookList)
	my.GET("/book", controller.UserBook)

	//类别
	category := r.Group("cate")
	category.GET("/list", controller.CategoryList)
	category.POST("/add", controller.CategoryAdd)
	category.PUT("/info", controller.CategoryUpdate)
	category.DELETE("/del", controller.CategoryDelete)

	// 图书
	book := r.Group("/letter")
	book.GET("/list", controller.BookList)
	book.POST("/add", controller.BookAdd)
	book.GET("/info", controller.BookGetDetail)
	book.PUT("/info", controller.BookEdit)
	book.DELETE("/info", controller.BookDelete)

	//用户信息
	user := r.Group("/user")
	user.GET("/list", controller.UserList)
	user.POST("/add", controller.UserAdd)
	user.DELETE("/info", controller.UserDelete)
	user.PUT("/info", controller.UserEditRoot)

	record := r.Group("/record")
	//借书
	record.POST("/borrow", controller.BorrowBook)
	//还书
	record.POST("/lend", controller.LendBook)
	record.GET("/list", controller.RecordList)

	r.Run()
}
