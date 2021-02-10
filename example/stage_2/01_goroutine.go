/*
 * @Descripttion: goroutine 协程 （轻量级的线程），协程是并发执行的
 * @Author: tacks321@qq.com
 * @Date: 2021-02-10 15:11:16
 * @LastEditTime: 2021-02-10 15:23:10
 */

package main

import "fmt"

// 函数
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	fmt.Println("============================goroutine协程================================")

	// 直接同步调用 （阻塞式调用的输出）
	f("direct")

	// ==========================goroutine协程= 两个协程交替出现。 Go 运行时是以 并发的方式运行协程的=====================
	// 启动协程调用
	go f("goroutine")

	// 匿名函数启用协程
	go func(msg string) {
		for i := 0; i < 3; i++ {
			fmt.Println(msg, ":", i)
		}
	}("going")

	// ==========================goroutine协程======================

	fmt.Scanln()
	fmt.Println("done")
}
