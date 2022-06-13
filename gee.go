package gee

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

// 某个结构体实现了接口内的所有函数 就认为实现了该接口
// 定义空结构体Engine，实现ServeHTTP方法

// Handler是一个接口，需要实现其中的ServeHTTP方法，用来解析URL并进行路由匹配
// 传入实现了ServeHTTP的实例后，所有HTTP请求都交由该实例来处理
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// request包含请求的具体信息
	// responseWriter可以构造请求的返回值

	switch req.URL.Path {
	case "/":
		// 不同的路由，不同的处理方法
		fmt.Fprintf(w, "URL PATH is %q/n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			// 请求头中的key-value信息
			fmt.Fprintf(w, "Header[%q] = %q/n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 not found URL %s/n", req.URL)
	}
}

//  启动路由
func run() {
	// 创建ServeHTTP的实例
	engine := new(Engine)
	// 启动web服务，绑定在9999端口
	// engine为我们实现的web框架的入口
	log.Fatal(http.ListenAndServe(":9999", engine))
}
