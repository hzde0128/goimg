package main

import (
	"fmt"
	"github.com/hzde0128/goimg/config"
	"github.com/hzde0128/goimg/router"
	"github.com/hzde0128/goimg/server"
)

func main() {

	// 初始化路由
	router.InitRoute()

	fmt.Printf("goimg v1.0.0\n  You can run goimg with env:\n  SERVER_PORT=port IMAGE_PATH=/path/to/img ./goimg\n\n")

	// 开始监听
	server.RunHTTP(config.HTTPAddr())
}
