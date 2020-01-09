package main

import (
	"fmt"
	"github.com/hzde0128/goimg/route"
	"github.com/hzde0128/goimg/server"
	"github.com/hzde0128/goimg/config"
)

func main() {

	// 初始化路由
	route.InitRoute()

	fmt.Printf("goimg v1.0.0\n  You can run goimg with env:\n  SERVER_PORT=port IMAGE_PATH=/path/to/img ./goimg\n")

	// 开始监听
	server.RunHttp(config.HttpAddr())
}
