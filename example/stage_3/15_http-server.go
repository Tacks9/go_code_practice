/*
 * @Descripttion:  http-server服务端
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 18:03:05
 * @LastEditTime: 2021-02-20 18:11:31
 */

package main

import (
	"fmt"
	"net/http"
)

// handler对象  handlers 是 `net/http` 服务器里面的一个基本概念。
// http.ResponseWriter 响应内容
// req 请求的信息
func hello(w http.ResponseWriter, req *http.Request) {
	// 输出到响应主体上
	fmt.Fprintf(w, "Hello\n")
	// 请求客户端的信息打印到终端上
	fmt.Println(req)
}

//
func headers(w http.ResponseWriter, req *http.Request) {

	// 这个 handler 稍微复杂一点，
	// 我们需要读取的 HTTP 请求 header 中的所有内容，并将他们输出至 response body。
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	fmt.Println("=============================【 服务端 】======================================")

	// 使用 `http.HandleFunc` 函数，可以方便的将我们的 handler 注册到服务器路由。
	// 它是 `net/http` 包中的默认路由，接受一个函数作为参数。
	// curl localhost:8090/hello
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 最后，我们用端口和处理程序调用 `ListenAndServe`。`nil` 告诉它使用我们刚刚设置的默认路由器。
	http.ListenAndServe(":8090", nil)
}

// 可以请求 http://localhost:8090/headers
// 可以请求 http://localhost:8090/hello
