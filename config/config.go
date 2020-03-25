package config

// goimg配置项

import (
	"os"
)

// HTTPAddr 监听端口
func HTTPAddr() string {

	// 从环境变量SERVER_PORT获取启动端口
	addr := os.Getenv("SERVER_PORT")

	if addr != "" {
		if addr[0] != ':' {
			addr = ":" + addr
		}
	} else {
		addr = ":8080"
	}
	return addr
}

// PathImg 图片存储目录
func PathImg() (path string) {

	// 从环境变量IMAGE_PATH获取图片存储路径
	path = os.Getenv("IMAGE_PATH")
	paths := []byte(path)
	if path != "" {
		if paths[len(path)-1] != '/' {
			path = path + "/"
		}
	}
	if path == "" {
		path = "/data"
	}
	return
}
