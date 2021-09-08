package main

import (
	"go-demo/initialize"
)

func main() {
	//初始化路由
	r := initialize.InitRouter()

	r.Run()
}
