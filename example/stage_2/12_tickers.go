/*
 * @Descripttion: 打点器 tickers
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 15:36:21
 * @LastEditTime: 2021-02-18 15:58:29
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=============================【 打点器 Tickers 】======================================")

	// 打点器： 在固定时间间隔内重复执行

	// 每隔200ms发送一次值
	ticker := time.NewTicker(time.Millisecond * 200)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 将在运行后 1000ms 停止这个打点器。
	time.Sleep(time.Millisecond * 1000)
	ticker.Stop()

	// 一旦打点器停止，那么就不能在通道里读出值
	fmt.Println("Ticker Stopped")
}
