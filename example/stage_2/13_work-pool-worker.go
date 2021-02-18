/*
 * @Descripttion: 线程池
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 16:07:02
 * @LastEditTime: 2021-02-18 16:45:36
 */
package main

import (
	"fmt"
	"time"
)

/**
worker()


id 	   任务号
jobs   接收数据
results发送数据
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	// 接收数据
	for j := range jobs {
		fmt.Println("worker", id, "Processing job", j)
		// 模拟耗时1s
		time.Sleep(time.Second)
		// 写入数据
		results <- j * 2
	}
}

func main() {
	fmt.Println("=============================【 线程池 work pool 】======================================")

	// 任务通道
	jobs := make(chan int, 100)
	// 结果通道
	results := make(chan int, 100)

	// 启动3个worker
	for i := 1; i <= 3; i++ {
		// 刚开始是阻塞状态，因为目前还没有jobs
		go worker(i, jobs, results)
	}

	// 发送9个jobs 然后关闭通道
	// 整个程序 处理所有的任务仅执行了 3s ，因为有3个worker并行，协程的强大
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后收集所有任务的返回值
	for k := 1; k <= 9; k++ {
		<-results
	}
}
