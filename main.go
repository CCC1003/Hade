package main

import (
	"Hade/Hin"
	"Hade/Hin/middleware"
	"net/http"
)

func main() {
	core := Hin.NewCore()
	core.Use(middleware.Recovery())
	registerRouter(core)
	server := &http.Server{
		//自定义核心处理函数
		Handler: core,
		//请求监听地址
		Addr: ":8888",
	}
	server.ListenAndServe()
}
