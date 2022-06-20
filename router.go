package gee

import "net/http"

// router 中 HandleFunc 究竟是什么呢
type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

func (router *router) addRouter(method string, pattern string, handleFunc HandleFunc) {
	key := method + "-" + pattern
	router.handlers[key] = handleFunc
}
func (router *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := router.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Req.URL)
	}
}
