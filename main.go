package main

import "go-demo/initialize"

func main() {
	//初始化logger
	initialize.InitLogger()
	//初始化config
	initialize.InitConfig()
	//初始化数据库
	initialize.InitDb()
	defer initialize.Db.Close()
	//初始化路由
	r := initialize.InitRouter()
	r.Run()
}
