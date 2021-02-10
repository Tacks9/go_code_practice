/*
 * @Descripttion: 通道的方向 direction
 * @Author: tacks321@qq.com
 * @Date: 2021-02-10 16:08:49
 * @LastEditTime: 2021-02-10 16:41:05
 */

package main

import "fmt"

// 发送通道数据 函数  （chan<- 发送）
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 接收通道数据 函数 （<-chan 接收）
func pong(pings <-chan string, pongs chan<- string) {
	// 从pings中接收到数据，存入msg
	msg := <-pings
	// 然后将msg数据发送给 pongs通道中
	pongs <- msg
}
func main() {
	fmt.Println("=============================【通道的方向 Channel Direction】======================================")

	// 创建两个通道
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// 向 pings通道中发送消息
	ping(pings, "push msg")
	// 将接收的信息再发送到 pongs
	pong(pings, pongs)

	fmt.Println(<-pongs)

}
