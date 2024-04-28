package main

import (
	"bookmanage/common"
	"bookmanage/router"
)

func main() {

	common.SetupDb()
	router.Run()

}
