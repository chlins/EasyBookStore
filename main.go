package main

import (
	"golang-web-testcase/router"
)

func main() {
	//加载路由
	router := router.Load()
	//加载模板
	router.LoadHTMLGlob("templates/*")
	router.Run(":8888")
}
