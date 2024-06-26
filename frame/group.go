package frame

import "log"

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // 中间件
	parent      *RouterGroup  // 父group
	engine      *Engine       // 全部group共用一个Engine
}

// Use
// @Description: 添加中间件
// @receiver group
// @param middlewares
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// Group
// @Description: 创建一个 RouterGroup
// @receiver group
// @param prefix
// @return *RouterGroup
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s\n", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
