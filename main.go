package main

import "go-demo/init"

func main() {
	//初始化logger
	init.InitLogger()
	//初始化config
	init.InitConfig()
	//初始化数据库
	init.InitDb()
	defer init.Db.Close()
	//初始化路由
	r := init.InitRouter()
	r.Run()
}
