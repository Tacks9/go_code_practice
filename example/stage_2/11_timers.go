/*
 * @Descripttion: 定时器 timers  time.NewTimer
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 15:22:47
 * @LastEditTime: 2021-02-18 15:34:21
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=============================【 定时器 timers 】======================================")

	// 定时器 ： 未来某一时刻的独立事件

	//========================================================================

	// 创建一个2s的定时器通道
	timer1 := time.NewTimer(time.Second * 2)

	// 定时器通道。 在失效之前将一直阻塞
	<-timer1.C
	fmt.Println("Timer1 Expired")

	// 当然，你这样看上面的这个例子好像跟 time.sleep 效果一样。
	// 所以如果你只是需要等待两秒，那你可以直接使用 time.sleep 而不是定时器

	//========================================================================
	timer2 := time.NewTimer(time.Second)
	// 协程
	go func() {
		<-timer2.C
		fmt.Println("Timer2 Expired!")
	}()

	// 定时器重要的作用，就是可以在定时器失效之前，取消这个定时器
	// 还没到时间，就提前结束定时器
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 Stopped")
	}
}
