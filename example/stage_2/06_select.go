/*
 * @Descripttion: 通道选择器 Channel Select
 * @Author: tacks321@qq.com
 * @Date: 2021-02-10 16:42:22
 * @LastEditTime: 2021-02-10 18:39:40
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=============================【通道选择器 Channel Select 】======================================")

	// 创建两个通道
	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在若干时间后接收一个值

	// 协程中，写入one
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()

	// 协程中，写入two
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//  `select` 关键字来同时等待这两个值
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("[Received]: ", msg1)
		case msg2 := <-c2:
			fmt.Println("[Received]: ", msg2)
		}
	}

}
