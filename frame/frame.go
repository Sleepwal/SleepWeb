package frame

import (
	"net/http"
)

// HandlerFunc 提供给用户，定义用于处理HTTP请求的方法
type HandlerFunc func(*Context)

// Engine 实现ServeHTTP接口
type Engine struct {
	router *router // 路由
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
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
	c := newContext(w, req)
	engine.router.handle(c)
}
