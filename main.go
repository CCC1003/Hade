package main

import (
	"Hade/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		//自定义核心处理函数
		Handler: framework.NewCore(),
		//请求监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
}
