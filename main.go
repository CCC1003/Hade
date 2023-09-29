package main

import (
	"Hade/framework"
	"Hade/framework/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	registerRouter(core)
	server := &http.Server{
		//自定义核心处理函数
		Handler: core,
		//请求监听地址
		Addr: ":8888",
	}

	go func() {
		server.ListenAndServe()
	}()
	//当前的goroutine等待信号量
	quit := make(chan os.Signal)
	//监控信号：SIGINT，SYSTERM，SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//这里会阻塞当前goroutine等待信号
	<-quit
	// 调用Server.Shutdown 结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
