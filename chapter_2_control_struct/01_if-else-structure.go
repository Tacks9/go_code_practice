// 2020-11-11 周一

package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("-------------------------------------this is if else")

	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	str := "tacks"
	if str == "" {
		fmt.Println("str is empty")
	} else {
		fmt.Println("str is not empty")
	}

	if runtime.GOOS == "windows" {
		fmt.Println("this is windows")
	} else { // Unix-like
		fmt.Println("this is unix")
	}

	fmt.Println("-------------------------------------if func")
	x := -10
	fmt.Println("this is %d", Abs(x))
	y := 2
	fmt.Println("this is %t", isGreater(x, y))

	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n") // first is 5 or greater
	}
	// 使用简短方式 := 声明的变量的作用域只存在于 if 结构中
	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n") // cond is not greater than 10
	}

}

/*
 返回一个整型数字的绝对值
*/
func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

/*
比较两个整型数字的大小
*/
func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

/*

if
	1、尽可能先满足的条件放在前面
	2、关键字 if 和 else 之后的左大括号 { 必须和关键字在同一行
	3、如果使用了 else-if 结构，则前段代码块的右大括号 } 必须和 else-if 关键字在同一行
	4、if条件中，不需要加小括号
	5、使用简短方式 := 声明的变量的作用域只存在于 if 结构中
*/
