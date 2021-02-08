/*
 * @Descripttion: go 数组 （具有相同类型的，固定长度，最长2GB，值类型）
 * @Author: tacks321@qq.com
 * @Date: 2021-02-07 17:56:30
 * @LastEditTime: 2021-02-08 18:31:59
 */

package main

import "fmt"

func main() {
	fmt.Println("========================arrays数组===============================")
	//【1】 一维数组声明   var ArrayName [Size]Type

	// 默认初始化为 int的零值. 也就是都是0 0
	var arr1 [5]int
	fmt.Println("[arr1] ", arr1)

	// 【2】采用 ArrayName[Index] = Value 进行赋值
	arr1[0] = 18
	arr1[1] = 24
	fmt.Println("set:", arr1)

	// 【3】采用 ArrayName[Index] 获取数组的值
	fmt.Println("get:arr[1] ", arr1[1])

	// 【4】声明并且初始化赋值   ArrayName := [Size]Type{...}
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("[arr2]: ", arr2)

	// 【5】声明多维数组
	var twoarr [2][4]int
	// 多维数组赋值
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			twoarr[i][j] = i + j
		}
	}
	fmt.Println("[twoarr]: ", twoarr)
	// 采用for range 遍历数组   for key := range ArrayName {}
	for i := range twoarr {
		fmt.Printf("Array at index %d is %d \n", i, twoarr[i])
	}

	// 【5】 字符串数组 ...符号指定的长度等于数组中元素的数量，当然...可以忽略
	stringArr := [...]string{"a", "b", "c", "d"} // 初始化数组
	// 采用 for range进行遍历数组
	for i := range stringArr {
		fmt.Print(stringArr[i], " ")
	}
	fmt.Println()

	// 【6】 数组内部分原始元素初始化
	var arrKeyVal = [5]string{2: "xiao-hong", 4: "xiao-lan"}
	fmt.Println(arrKeyVal)

	// 【7】使用内置函数“len” 返回数组的长度
	fmt.Println("[arrKeyVal] len:", len(arrKeyVal))

	// 【8】数组发生内存拷贝
	arrCopy1 := [5]int{1, 2, 3, 4, 5}
	// 内存拷贝
	arrCopy2 := arrCopy1
	// 修改后
	arrCopy2[0] = 0
	fmt.Println(arrCopy1)
	fmt.Println(arrCopy2)

}
