package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler // 从core这边设置的中间件
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

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = middlewares
}

// Get 匹配GET 方法, 增加路由规则
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	// 将core的middleware 和 handlers结合起来
	allhandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allhandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Post 匹配POST 方法, 增加路由规则
func (c *Core) Post(url string, handlers ...ControllerHandler) {
	// 将core的middleware 和 handlers结合起来
	allhandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allhandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Put 匹配PUT 方法, 增加路由规则
func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allhandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allhandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Delete 匹配DELETE 方法, 增加路由规则
func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allhandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allhandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// FindRouteByRequest 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) *node {
	//uri 和 method全部转换为大写，保证大小写不敏感
	path := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	//查找第一层
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.root.matchNode(path)
	}
	return nil
}
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	//封装自定义context
	ctx := NewContext(request, response)
	//寻找路由
	node := c.FindRouteByRequest(request)
	if node == nil {
		//如果没有找到，这里打印日志
		ctx.SetStatus(404).Json("not found")
		return
	}
	ctx.SetHandlers(node.handlers)

	// 设置路由参数
	params := node.parseParamsFromEndNode(request.URL.Path)
	ctx.SetParams(params)

	//调用next函数，如果返回err 代表存在内部错误，返回500状态码
	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("inner error")
		return
	}
}
