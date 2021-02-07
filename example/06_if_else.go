/*
 * @Descripttion: if-else分支 条件判断
 * @Author: tacks321@qq.com
 * @Date: 2021-02-07 17:00:05
 * @LastEditTime: 2021-02-07 17:05:57
 */

package main

import "fmt"

func main() {
	fmt.Println("===================if-else=====================")

	// 只有if分支
	if 8%2 == 0 {
		fmt.Println("8 is even")
	}

	// if else 两个分支
	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	// if分支上可以加上声明语句
	if num := 9; num < 0 {
		fmt.Println(num, "num<0")
	} else if num < 10 {
		fmt.Println(num, "num<10")
	} else {
		fmt.Println(num, "num>=10")
	}

}
