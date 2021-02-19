/*
 * @Descripttion: 协程状态 stateful
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 18:31:48
 * @LastEditTime: 2021-02-19 09:54:53
 */
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 在这个例子中，state 将被一个单独的 Go 协程拥有。这就
// 能够保证数据在并行读取时不会混乱。为了对 state 进行
// 读取或者写入，其他的 Go 协程将发送一条数据到拥有的 Go
// 协程中，然后接收对应的回复。结构体 `readOp` 和 `writeOp`
// 封装这些请求，并且是拥有 Go 协程响应的一个方式。
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// 记录操作次数
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// 读请求 / 写请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 这个就是拥有 `state` 的那个 Go 协程
	go func() {
		// map
		var state = make(map[int]int)
		// 反复响应
		for {
			select {
			// 读请求
			case read := <-reads:
				read.resp <- state[read.key]
				// 写请求
			case write := <-writes:
				state[write.key] = write.val
				// 响应值设置 true 表示成功
				write.resp <- true
			}
		}
	}()

	// 启动 100 个 协程 处理读请求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// 每个读请求都构造一个 readOp
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				// 写入通道
				reads <- read
				// 阻塞响应 接收到通道响应
				<-read.resp
				// 计数
				atomic.AddUint64(&readOps, 1)
				// 休眠
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 用相同的方法启动 10 个写操作。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 让 Go 协程们跑 1s。
	time.Sleep(time.Second)

	// 读请求
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	// 写请求
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
