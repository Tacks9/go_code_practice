/*
 * @Descripttion: 速率限制 rate-limiting
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 17:11:26
 * @LastEditTime: 2021-02-18 17:27:49
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=============================【 速率限制 rate-limiting 】======================================")

	// 通过Go协程，通道，打点器 来完成速率限制

	// ==================================================================
	fmt.Println("【第一批匀速处理接收请求】")
	// 创建一个通道
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		// 发送内容
		requests <- i
	}
	// 关闭通道
	close(requests)

	// 速率限制管理器
	// 将limiter通道200ms接收一个值
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		// 阻塞中
		<-limiter
		// 每200ms接受一个值
		fmt.Println("request", req, time.Now())
	}

	// ==================================================================
	fmt.Println("【第二批 脉冲型速率限制 接收请求】")
	burstyLimiter := make(chan time.Time, 3)

	// 三次
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// 每200ms加入一个新的值
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	// 5个数据
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		// 脉冲型速率缓冲
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
