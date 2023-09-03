# Hade

基于Go语言从0开发一个基础Web框架

### 一、使用Golang net-http标准库搭建Server

net/http标准库创建服务的主流程：

1. 第一层：http.ListenAndServe 本质是通过创建一个 Server 数据结构，调用server.ListenAndServe 对外提供服务，这一层完全是比较简单的封装，目的是，将Server 结构创建服务的方法 ListenAndServe ，直接作为库函数对外提供，增加库的易用性。
2. 第二层：创建服务的方法 ListenAndServe 先定义了监听信息 net.Listen监听地址，然后调用 Serve 函数。
3. 第三层：Serve 函数中，用了一个 for 循环，通过 l.Accept不断接收从客户端传进来的请求连接。当接收到了一个新的请求连接的时候，通过 srv.NewConn创建了一个连接结构（http.conn），并创建一个 Goroutine 为这个请求连接对应服务（c.serve）。
4. 第四层：c.serve函数先判断本次 HTTP 请求是否需要升级为 HTTPs，接着创建读文本的 reader 和写文本的 buffer，再进一步读取本次请求数据
5. 第五层：调用最关键的方法 serverHandler{c.server}.ServeHTTP(w, w.req) ，来处理这次请求。
6. 第六层：如果入口服务 server 结构已经设置了 Handler，就调用这个 Handler 来处理此次请求，反之则使用库自带的 DefaultServerMux。
7. 第七层：DefaultServerMux 是使用 map 结构来存储和查找路由规则。DefaultServeMux.Handle 是一个非常简单的 map 实现，key 是路径（pattern）， value 是这个 pattern 对应的处理函数（handler）。它是通过 mux.match(path) 寻找对应 Handler，也就是从 DefaultServeMux 内部的 map 中直接根据 key 寻找到 value的。

### 二、Context:上下文请求控制器，实现链条信息传递和共享

从主流程中（第三层关键结论），HTTP 服务会为每个请求创建一个 Goroutine 进行服务处理。在服务处理的过程中，有可能就在本地执行业务逻辑，也有可能再去下游服务获取数据。会形成一个标准的树形逻辑链条。

在整个树形逻辑链条中，用上下文控制器Context，实现每个节点的信息传递和共享。

### 三、Route:路由规则（使用trie树结构）:


基本需求从简单到复杂，为下面四点
1. HTTP方法匹配
2. 静态路由匹配
3. 批量通用前缀（接口替代结构定义解耦合）
4. 动态路由匹配

trie 树不同于二叉树，它是多叉的树形结构，根节点一般是空字符串，而叶子节点保存的通常是字符串，一个节点的所有子孙节点都有相同的字符串前缀。

这个 trie 树是按照路由地址的每个段 (segment) 来切分的，每个 segment 在 trie 树中都能找到对应节点，每个节点保存一个 segment。树中，每个叶子节点都代表一个 URI，对于中间节点来说，有的中间节点代表一个 URI（比如 /subject/name），而有的中间节点并不是一个 URI（因为没有路由规则对应这个 URI）。
 
- 定义树和节点的数据结构
- 编写函数：增加路由规则（AddRouter）
- 编写函数：查找路由（FindHandler）
- 将增加路由规则和查找路由规则添加到框架中

### 四、middleware:使用pipeline模式设计框架中间件

装饰器模式：核心处理模块的外层增加一个又一个的装饰，类似洋葱。

问题：

1. 中间件是循环嵌套的，当有多个中间件的时候，整个嵌套长度就会非常长，可读性差。
2. 只能为当个业务控制器设置中间件，不能批量设置。

使用pipeline思想设计中间件

一层层嵌套不好用，如果我们将每个核心控制器所需要的中间件，使用一个数组链接 （Chain）起来，形成一条流水线（Pipeline），就能完美解决这两个问题了。

将每个中间件构造出来的ControllerHandler和最终的业务逻辑的ControllerHandler结合在一起，成为一个ControllerHandler数组，也就是控制器链。

引入了 pipeline 的思想，将所有中间件做成一个链条，通过这个链条的调用，来实现中间件机制。

### 五、request&response:使用接口定义，response定义函数实现链式调用

封装： 定义接口让封装更明确（对于完成的功能模块，先定义接口，再具体实现）

定义一个清晰的、包含若干个方法的接口，可以让使用者非常清楚：这个功能模块提供了哪些函数，哪些函数是需要的，哪些是不需要的，查找方便。

其次，定义接口可以"实现解耦"。使用接口作为参数、返回值，能够让使用者在写具体函数的时候，有不同的实现；而且在不同实现中，只需要做到接口一致，就能很简单进行替换，而不用修改使用方的任何代码。


**封装请求信息（request）：**

1. 参数信息
- URL携带参数
- 路由解析后带的参数
- Body内带的参数
2. Header信息

- cookie信息
- 基础信息（请求地址、请求方法、请求ip、请求域名）
- 其他信息

**封装响应信息（response）：**

这里实现的很多方法的返回值使用 IResponse 接口本身， 这个设计能允许使用方进行链式调用。链式调用的好处是，能很大提升代码的阅读性。

1. Header设置

- cookie设置
- 状态码设置（重定向状态码、正常状态码、其他状态码）
- 其他设置

2. Body设置

- JSON格式
- XML格式
- JSONP格式
- HTML格式


### 六、使用os/signal标准库和server.Shutdown实现框架等待所有请求逻辑处理结束后关闭服务