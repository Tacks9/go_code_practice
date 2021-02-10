/*
 * @Descripttion: channel管道
 * @Author: tacks321@qq.com
 * @Date: 2021-02-10 15:40:07
 * @LastEditTime: 2021-02-10 15:44:58
 */

package main

import "fmt"

func main() {
	fmt.Println("=====================channel 管道=======================")

	// 默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕

	// 创建一个通道类型,传值的类型为string
	messages := make(chan string)

	// 发送方【 channel <- 内容】
	// 在匿名函数启动的一个协程中，发送消息到通道中
	go func() {
		messages <- "ping1"
	}()

	// 接收方【 <-channel】
	// 然后采用  <- 接收通道的值并打印
	msg := <-messages
	fmt.Println(msg)
}
