package main

import (
	"goimg/route"
	"goimg/server"
	"goimg/config"
)

func main() {

	// 初始化路由
	route.InitRoute()

	// 开始监听
	server.RunHttp(config.HttpAddr())
}
