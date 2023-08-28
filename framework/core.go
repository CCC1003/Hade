package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router map[string]*Tree
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	// 初始化路由
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}

	return &Core{router: router}
}

// Get 匹配GET 方法, 增加路由规则
func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Post 匹配POST 方法, 增加路由规则
func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Put 匹配POST 方法, 增加路由规则
func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Delete 匹配POST 方法, 增加路由规则
func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// FindRouteByRequest 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	//uri 和 method全部转换为大写，保证大小写不敏感
	path := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	//查找第一层
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(path)
	}
	return nil
}
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	//封装自定义context
	ctx := NewContext(request, response)
	//寻找路由
	router := c.FindRouteByRequest(request)
	if router == nil {
		//如果没有找到，这里打印日志
		ctx.Json(404, "not found")
		return
	}
	//调用路由函数，如果返回err 代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
