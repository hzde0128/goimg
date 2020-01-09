package main

import (
	"github.com/hzde0128/goimg/route"
	"github.com/hzde0128/goimg/server"
	"github.com/hzde0128/goimg/config"
)

func main() {

	// 初始化路由
	route.InitRoute()

	// 开始监听
	server.RunHttp(config.HttpAddr())
}
