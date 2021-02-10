/*
 * @Descripttion: 变量类型，这里暂时只声明几种简单的变量类型，如int\string\bool
 * @Author: tacks321@qq.com
 * @Date: 2021-02-07 10:39:28
 * @LastEditTime: 2021-02-07 11:00:03
 */

package main

import "fmt"

func main() {
	fmt.Println("==========================变量类型=============================")
	fmt.Println("变量的声明与赋值，var identifier [type] = value")

	// 声明一个 字符串变量
	var str1 string
	// 变量赋值
	str1 = "Tacks"
	fmt.Println(str1)

	// 声明并且赋值
	var num1 int = 9
	fmt.Println(num1)

	// 短变量运算符
	is_ok := true
	fmt.Println(is_ok)

	// 多变量平行赋值
	a, b := "A", "B"
	fmt.Println(a, b)

	// 变量类型 自动推断
	var do = "做"
	fmt.Println(do)

	// 变量初始化后，go会自动赋值给默认值,例如int的默认值为0
	var defaultNum int
	fmt.Println(defaultNum)

}
