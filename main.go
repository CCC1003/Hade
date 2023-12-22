package main

import (
	"Hade/Horm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `horm:"primary key"`
	Age  int
}

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
	engine, _ := Horm.NewEngine("mysql", "root:123456@tcp(8.130.85.112:3306)/Horm")
	defer engine.Close()
	s := engine.NewSession().Model(&User{})
	_ = s.CreateTable()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "cong", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success,%d affected\n", count)
	_ = s.DropTable()
	if s.HasTable() {
		fmt.Println(s.HasTable())
	}
	//_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()

}
