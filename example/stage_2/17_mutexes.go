/*
 * @Descripttion: 互斥锁 sync.Mutex mutex unMutex
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 18:03:44
 * @LastEditTime: 2021-02-18 18:30:34
 */
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=============================【 互斥锁 mutexes 】======================================")

	// map
	var state = make(map[int]int)

	// 这里 mutex 将同步对 state的访问
	var mutex = &sync.Mutex{}

	// 记录我们对state的操作数
	var ops int64 = 0

	// 10个协程模拟写操作
	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100) // 随机生成值

				mutex.Lock()
				state[key] = val
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()

	}

	// 100个协程重复读取 state
	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				// 每次循环读取，我们使用一个键来进行访问
				// Lock() 这个 mutex 来确保对 state 的独占访问， 读取选定的值
				// Unlock() 这个 mutex 并 ops 值 +1
				key := rand.Intn(5)

				// 上锁在再进行操作
				mutex.Lock()
				total += state[key] // 随机获取值
				mutex.Unlock()

				// 计数器
				atomic.AddInt64(&ops, 1)

				// 释放
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	// 获取最后的操作计数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	// 加个锁看看
	mutex.Lock()
	fmt.Println("State：", state)
	mutex.Unlock()

}
