/*
 * @Descripttion: closing-of-channel 通道关闭
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 14:58:36
 * @LastEditTime: 2021-02-18 15:07:37
 */
package main

import "fmt"

func main() {
	fmt.Println("=============================【 通道关闭 close channel 】======================================")

	jobs := make(chan int, 5)
	done := make(chan bool)

	// 开启一个协程
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received Job", j)
			} else {
				fmt.Println("Received All Jobs")
				// 阻塞中
				done <- true
				// 结束
				return
			}
		}
	}()

	// 发送三个任务到jobs中
	for j := 1; j <= 3; j++ {
		// 写入通道
		jobs <- j
		fmt.Println("Sent job", j)
	}

	// 关闭通道,不能再写入了
	close(jobs)

	// 任务已经发送完
	fmt.Println("Sent All Jobs")

	// 通道在收到 true 之前一直阻塞
	<-done
}
