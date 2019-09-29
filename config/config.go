package config

import (
	"os"
)

func HttpAddr() string {

	// 从环境变量SERVER_PORT获取启动端口
	addr := os.Getenv("SERVER_PORT")

	if addr == "" {
		addr = ":8080"
	}
	addrs := []byte(addr)
	if addr != "" {
		if addrs[0] != ':' {
			addr = ":" + addr
		}
	}
	return addr
}

func PathImg() string {

	// 从环境变量IMAGE_PATH获取图片存储路径
	path := os.Getenv("IMAGE_PATH")
	paths := []byte(path)
	if path != "" {
		if paths[len(path)-1] != '/' {
			path = path + "/"
		}
	}
	if path == "" {
		path = "/data/"
	}

	return path
}
