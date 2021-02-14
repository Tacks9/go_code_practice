/*
 * @Descripttion:超时处理 Timeout
 * @Author: tacks321@qq.com
 * @Date: 2021-02-14 16:28:02
 * @LastEditTime: 2021-02-14 16:48:08
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=============================【超时处理 timeout】======================================")

	// 定义一个通道用来返回协程的结果
	chan1 := make(chan string, 1)

	go func() {
		// 休眠 1 s
		time.Sleep(time.Second * 2)
		// 写入通道
		chan1 <- "result 1"
	}()

	select {
	// 等待结果
	case res := <-chan1:
		fmt.Println(res)
		// 等待超时
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1 ")
	}

	chan2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- "result * 2"
	}()

	// 使用这个 select 超时方式，需要使用通道传递结果 chan2
	select {
	case res := <-chan2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 3")
	}

}
