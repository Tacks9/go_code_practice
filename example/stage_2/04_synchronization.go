/*
 * @Descripttion: Channel Synchronization 通道同步
 * @Author: tacks321@qq.com
 * @Date: 2021-02-10 15:55:22
 * @LastEditTime: 2021-02-10 16:03:36
 */

package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working ...", time.Second*3)
	time.Sleep(time.Second * 3)
	fmt.Println("done !")
	time.Sleep(time.Second * 3)

	// 发送一个值表示我们结束了
	done <- true
}
func main() {
	fmt.Println("============================Channel Synchronization =======================")
	// 利用通道来同步 Go 协程间的执行状态

	// 创建一个通道
	done := make(chan bool, 1)
	go worker(done)

	// 通道在收到 work发送的 true 之前一直阻塞
	<-done

}
