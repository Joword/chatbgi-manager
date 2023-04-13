package src

import (
	"fmt"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	version := "1.0.0"

	fmt.Printf(`
		欢迎使用 ChatBGI-Vue-Manager
		当前版本:` + version + `
		默认自动化文档地址:http://127.0.0.1/swagger/index.html
		默认前端文件运行地址:http://127.0.0.1:9090
	`)
}
