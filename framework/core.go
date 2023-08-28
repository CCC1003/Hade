package framework

import (
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router map[string]map[string]ControllerHandler
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	//二级map
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}
	//将二级map写入一级map
	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter

	return &Core{router: router}
}
func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}
func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}
func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}
func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}

// FindRouteByRequest 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	//uri 和 method全部转换为大写，保证大小写不敏感
	path := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	upperUri := strings.ToUpper(path)

	//查找第一层
	if methodHandlers, ok := c.router[upperMethod]; ok {
		//查找第二层
		if handler, ok := methodHandlers[upperUri]; ok {
			return handler
		}
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
