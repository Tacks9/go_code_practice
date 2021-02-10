/*
 * @Descripttion: 常量类型
 * @Author: tacks321@qq.com
 * @Date: 2021-02-07 11:02:51
 * @LastEditTime: 2021-02-07 11:19:15
 */

package main

import "fmt"

// 定义常量
const Name string = "Tacks"

func main() {

	fmt.Println("=======================常量========================")
	fmt.Println("定义格式 const identiferName [type] = value,例如 const PI = 3.1415926")

	fmt.Println(Name)

	// 常量数值 这个时候具体不太清楚是什么类型
	const age = 18
	// 数值型常量没有确定的类型，直到被给定 ，可以进行一次显示转化
	fmt.Println(int64(age))

	// 常量类型，必须在编译的时候就知道确定的值
	const nameLen = len(Name)
	fmt.Println(nameLen)
}
