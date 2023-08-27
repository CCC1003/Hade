package main

import (
	"Hade/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		//自定义核心处理函数
		Handler: core,
		//请求监听地址
		Addr: ":8888",
	}
	server.ListenAndServe()
}
