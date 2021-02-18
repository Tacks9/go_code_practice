/*
 * @Descripttion: waitgroups 协程
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 17:00:48
 * @LastEditTime: 2021-02-18 17:10:25
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

/**
worker() 任务函数

WaitGroup 必须通过指针传递给函数
*/

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d Starting\n", id)

	// 睡眠一秒钟 模拟耗时
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)

	// 通知 waitGroup 当前协程工作已经完成
	wg.Done()
}

func main() {
	fmt.Println("=============================【 协程 waitgroups 】======================================")

	// 用于等待所有开启的协程
	var wg sync.WaitGroup

	// 开启协程
	for i := 1; i <= 5; i++ {
		// 增加计数器
		wg.Add(1)
		// 执行任务
		go worker(i, &wg)
	}

	// 阻塞直到waitGroup计数器恢复到0，也就是所有的协程工作完成
	wg.Wait()
}
