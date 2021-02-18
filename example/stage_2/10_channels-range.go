/*
 * @Descripttion: 通道遍历  range迭代 注意close关闭
 * @Author: tacks321@qq.com
 * @Date: 2021-02-18 15:09:50
 * @LastEditTime: 2021-02-18 15:20:27
 */

package main

import "fmt"

func main() {
	fmt.Println("=============================【 通道遍历 channel range 】======================================")

	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"

	// 注意提前关闭通道,如果没有关闭，最后仍然会在阻塞中
	close(queue)

	// 关闭后的通道,仍然可以利用range进行迭代获取
	for elem := range queue {
		fmt.Println(elem)
	}

	// 等到迭代获取之后, queue通道就会进行关闭

}
