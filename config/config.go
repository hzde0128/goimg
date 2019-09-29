package config

import (
	"os"
)

func HttpAddr() string {

	// 从环境变量SERVER_PORT获取启动端口
	addr := os.Getenv("SERVER_PORT")
	addrs := []byte(addr)
	if addrs[0] != ':' {
		addr = ":" + addr
	}
	if addr == "" {
		addr = ":8080"
	}

	return addr
}

func PathImg() string {

	// 从环境变量IMAGE_PATH获取图片存储路径
	path := os.Getenv("IMAGE_PATH")
	paths := []byte(path)
	if paths[len(path)-1] != '/' {
		path = path + "/"
	}
	if path == "" {
		path = "img/"
	}

	return path
}
