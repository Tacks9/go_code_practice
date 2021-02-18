/*
 * @Descripttion: 非阻塞通道 non blocking
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 10:31:02
 * @LastEditTime: 2021-02-18 14:56:12
 */

package main

import "fmt"

func main() {
	fmt.Println("=============================【 非阻塞通道 non blocking】======================================")

	fmt.Println("[非阻塞接收]")
	messages := make(chan string)
	select {
	case msg := <-messages:
		fmt.Println("Receive Message", msg)
	default:
		fmt.Println("no message received")
	}

	fmt.Println("[非阻塞发送]")
	msg := "hello"
	select {
	case messages <- msg:
		fmt.Println("Set Message", msg)
	default:
		fmt.Println("no message sent")

	}

	fmt.Println("[非阻塞多路]")
	signals := make(chan bool)
	select {

	case msg := <-messages:
		fmt.Println("Received Message", msg)
	case sig := <-signals:
		fmt.Println("Received Signal", sig)
	default:
		fmt.Println("no activity")

	}

}
