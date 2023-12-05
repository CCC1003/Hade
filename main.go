package main

import (
	"Hade/Horm"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//hin
	//core := Hin.NewCore()
	//core.Use(middleware.Recovery())
	//registerRouter(core)
	//server := &http.Server{
	//	//自定义核心处理函数
	//	Handler: core,
	//	//请求监听地址
	//	Addr: ":8888",
	//}
	//server.ListenAndServe()

	//horm
	engine, _ := Horm.NewEngine("mysql", "")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success,%d affected\n", count)
}
