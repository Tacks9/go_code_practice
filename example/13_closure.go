/*
 * @Descripttion: closure 函数闭包
 * @Author: tacks321@qq.com
 * @Date: 2021-02-09 16:36:57
 * @LastEditTime: 2021-02-09 16:51:53
 */

package main

import "fmt"

// 闭包
func intSeq() func() int {
	// 对i变量进行隐藏
	i := 0
	// 返回匿名函数
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("===============closure闭包====================")

	nextInt := intSeq()

	// 通过多次调用 nextInt
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}
