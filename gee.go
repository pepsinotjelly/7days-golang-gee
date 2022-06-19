package gee

import (
	"fmt"
	"net/http"
)

// HandleFunc 定义使用的Handler的接口
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine 某个结构体实现了接口内的所有函数 就认为实现了该接口
// 定义结构体Engine，实现ServeHTTP方法
type Engine struct {
	router map[string]HandleFunc
}

// New 定义Engine的构造器
func New() *Engine {
	// 这个make函数是做啥的嘞？好像是创建用的类似new对吗
	// make专门创建 slice 哈希表 channel 这些复合类型的变量
	return &Engine{router: make(map[string]HandleFunc)}
}

func (engine *Engine) addRouter(method string, pattern string, handleFunc HandleFunc) {
	key := method + "-" + pattern
	engine.router[key] = handleFunc
}

func (engine *Engine) GET(pattern string, handleFunc HandleFunc) {
	engine.addRouter("GET", pattern, handleFunc)
}

func (engine *Engine) POST(pattern string, handleFunc HandleFunc) {
	engine.addRouter("POST", pattern, handleFunc)
}

// Run 启动 http serve
func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// Handler是一个接口，需要实现其中的ServeHTTP方法，用来解析URL并进行路由匹配
// 传入实现了ServeHTTP的实例后，所有HTTP请求都交由该实例来处理
//func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	// request包含请求的具体信息
//	// responseWriter可以构造请求的返回值
//
//	switch req.URL.Path {
//	case "/":
//		// 不同的路由，不同的处理方法
//		fmt.Fprintf(w, "URL PATH is %q/n", req.URL.Path)
//	case "/hello":
//		for k, v := range req.Header {
//			// 请求头中的key-value信息
//			fmt.Fprintf(w, "Header[%q] = %q/n", k, v)
//		}
//	default:
//		fmt.Fprintf(w, "404 not found URL %s/n", req.URL)
//	}
//}

////  启动路由
//func run() {
//	// 创建ServeHTTP的实例
//	engine := new(Engine)
//	// 启动web服务，绑定在9999端口
//	// engine为我们实现的web框架的入口
//	log.Fatal(http.ListenAndServe(":9999", engine))
//}
