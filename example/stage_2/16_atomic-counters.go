/*
 * @Descripttion: 原子计数器 "sync/atomic"
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 17:50:23
 * @LastEditTime: 2021-02-18 18:02:09
 */

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=============================【 原子计数器 】======================================")

	// 计数器
	var ops uint64 = 0

	// 启动50个协程
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 计数器加一
				atomic.AddUint64(&ops, 1)
				// 允许其他的go协程运行
				runtime.Gosched()
			}
		}()
	}

	// 等待1s，让ops的自加操作执行一会
	time.Sleep(time.Second)

	// 拷贝一下 为了安全的使用 在协程执行期间
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)
}
