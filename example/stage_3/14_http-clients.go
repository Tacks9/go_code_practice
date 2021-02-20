/*
 * @Descripttion: http-clients 客户端 （net/Http）
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 17:54:29
 * @LastEditTime: 2021-02-20 18:01:32
 */

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("=============================【 HTTP 客户端 】======================================")
	// 向服务端发送一个get请求
	resp, err := http.Get("http://baidu.com")
	if err != nil {
		panic(err)
	}
	// 记得及时关闭
	defer resp.Body.Close()

	// 打印Http Response 状态
	fmt.Println("Response Status:", resp.Status)

	// 打印响应主体 5 行
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	// 错误处理
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
