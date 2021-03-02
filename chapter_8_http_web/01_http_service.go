/*
 * @Descripttion: 模拟一个简单的http服务
 * @Author: tacks321@qq.com
 * @Date: 2020-12-01 18:03:16
 * @LastEditTime: 2021-03-02 15:48:18
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// 逻辑函数
func sayHelloName(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()   // 解析参数（默认不解析）
	fmt.Println(r.Form) // 服务端的信息

	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["name"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("Val:", strings.Join(v, ""))
	}
	// 输出内容到客户端浏览器
	_, _ = fmt.Fprintf(w, "HelloWorld!")

}
func main() {
	fmt.Println("[Http Service]")

	// 设置路由   注册了请求 / 的路由规则，当请求 uri 为 "/"，路由就会转到函数 sayHelloName
	http.HandleFunc("/", sayHelloName)
	// 设置监听端口
	err := http.ListenAndServe(":9000", nil)
	/*
			主要流程
			1、初始化一个 server 对象。
			2、然后调用net.Listen调用TCP协议搭建服务，进而监控对应的端口。
			3、然后底层调用srv.Serve(net.Listen)进行处理客户端请求信息。
				  函数中通过Listen接收请求Accept
		  		  然后创建一个conn=srv.newConn()
						客户端的每次请求都会创建一个 Conn，这个 Conn 里面保存了该次请求的信息，然后再传递到对应的 handler
		          最后单独go c.serve(connCtx)，也就是把这个请求的数据丢过去服务，每一个新的请求都是一个协程，互不影响。


	*/
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
【ServeMux】

type ServeMux struct {
	mu    sync.RWMutex		  // go 为了实现高并发高性能，采用 goroutines来处理conn的读写，这里需要采用一个锁机制
	m     map[string]muxEntry // 注册路由表达式 路由规则，一个string对应一个muxEntry实体
	// 用户请求的 URL 和路由器里面存储的 map 去匹配的，当匹配到之后返回存储的 handler，调用这个 handler 的 ServeHTTP 接口就可以执行到相应的函数
	es    []muxEntry // 实体的切片
	hosts bool       // 任意规则中是否包含host
}

type muxEntry struct {
	h       Handler  // 路由表达式对应不同的handler
	pattern string   // 匹配的字符串
}

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)  // 路由实现器
}

// 注册路由规则
type HandlerFunc func(ResponseWriter, *Request)

// 路由分发
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	// 如果是 * 关闭连接
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	// 不然就调用handler来处理路由
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}


*/
