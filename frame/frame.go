package frame

import (
	"fmt"
	"net/http"
)

// HandlerFunc 提供给用户，定义用于处理HTTP请求的方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现ServeHTTP接口
type Engine struct {
	router map[string]HandlerFunc // 路由
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET
// @Description: 注册GET请求的处理函数
// @receiver engine
// @param pattern
// @param handler
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST
// @Description: 注册POST请求的处理函数
// @receiver engine
// @param pattern
// @param handler
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP
// @Description: 实现ServeHTTP接口
// @receiver engine
// @param w
// @param req
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req) // 调用处理函数
	} else {
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
