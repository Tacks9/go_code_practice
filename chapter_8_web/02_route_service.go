/*
 * @Descripttion: 模拟一个简单的路由
 * @Author: tacks321@qq.com
 * @Date: 2021-03-02 15:04:03
 * @LastEditTime: 2021-03-02 15:48:00
 */
package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

// 路由分发逻辑
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

// handler逻辑处理
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 写入到浏览器中
	fmt.Fprintf(w, "Hello MyRoute!")
}

func main() {
	fmt.Println("【整一个简单的路由器】")
	mux := &MyMux{}
	http.ListenAndServe(":9000", mux)

}

/*
【整个流程】
	1、首先调用Http.HandleFunc
		调用默认的 DefaultServeMux.HandleFunc(pattern, handler)
		调用 mux.Handle(pattern, HandlerFunc(handler))
		增加 handler 路由规则 mux.m = make(map[string]muxEntry)
	2、然后调用http.ListenAndServe(":9000", mux)
		实例化 server := &Server{Addr: addr, Handler: handler}
		调用 server.ListenAndServe()
		调用 监听端口 ln, err := net.Listen("tcp", addr)
		调用 服务 把监听的端口的ln连接传入 srv.Serve(ln)
		启动一个 for 循环，在循环体中 Accept 请求
		对每一个实例化连接Conn 都会重启一个协程处理  进行服务 go c.serve ()
		在每一个请求都要读取内容 w, err := c.readRequest(ctx)
		然后调用 serverHandler{c.server}.ServeHTTP(w, w.req)
		然后判断是否有注册 ，有的话就设置handler ; 如果没有采用默认的 handler = DefaultServeMux
		然后进入handler调用ServeHttp
		进入路由选择 handler
			A 判断是否有路由能满足这个 request（循环遍历 ServeMux 的 muxEntry）
			B 如果有路由满足，调用这个路由 handler 的 ServeHTTP
			C 如果没有路由满足，调用 NotFoundHandler 的 ServeHTTP

*/
