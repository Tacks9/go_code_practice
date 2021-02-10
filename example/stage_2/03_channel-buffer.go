/*
 * @Descripttion: 通道缓冲 buffer
 * @Author: tacks321@qq.com
 * @Date: 2021-02-10 15:47:59
 * @LastEditTime: 2021-02-10 15:52:10
 */

package main

import "fmt"

func main() {
	// 默认情况下，通道是 无缓冲 的，这意味着只有对应的接收（<- chan） 通道准备好接收时，才允许进行发送（chan <-）。
	// 可缓存通道 允许在没有对应接收方的情况下，缓存限定数量的值。
	// 有点像缓冲队列，先攒一波，然后再统一进行接收。

	fmt.Println("=============================channel buffer 通道缓冲========================")
	// make一个字符串通道，最多缓存2个值
	messages := make(chan string, 2)

	// 由于通道可以缓冲，所以我们可以先将值发送到通道中
	messages <- "buffered"
	messages <- "channel"

	// 这个时候接收方 再进行接收
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
