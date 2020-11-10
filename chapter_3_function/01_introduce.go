// 2020-11-10 周二
package main

import (
	"fmt"
)

func main() {

	fmt.Println("-------------------------------------introduce function")

	fmt.Println("In main before calling greeting")
	greeting()
	fmt.Println("In main after calling greeting")

	fmt.Println("-------------------------------------函数可以将 其他函数调用 作为 本身函数的参数")

	a := 10
	b := 90
	bothNum := 9
	fmt.Println(sumTwo(sumThree(a, b, bothNum)))

}

// 打印 greeting函数
func greeting() {
	fmt.Println("In greeting: Hi!!!!!")
}

func sumTwo(a, b int) int {
	return a + b
}

func sumThree(a, b, bothNum int) (int, int) {
	return a + bothNum, b + bothNum
}

/*
1、函数是基本的代码块
2、Go是编译型语言，函数编写的顺序是无关紧要的，最好把 main() 函数放在最前面
3、主要目的：将需要多行代码的复杂问题拆解成一系列的简单任务，这就是函数的意义，而且函数的编写有利于代码重用
4、DRY（Don’t Repeat Yourself），特定任务的代码只能在程序中出现一次
5、return语句：可以携带0个或者一个或者多个参数返回(多值返回是 Go 的一大特性)，简单的return也可以终止for循环或者一个协程
6、三种类型的函数（普通函数、匿名函数、方法）
7、main()、init()函数没有参数和返回值
8、函数签名：函数参数、返回值、函数类型
9、其他包调用函数，pack1.Function(arg1,arg2...)
10、函数可以将 其他函数调用 作为 本身函数的参数。只要这个被调用函数的返回值个数、返回值类型和返回值的顺序与调用函数所需求的实参是一致的
11、在 Go 里面函数重载是不被允许的

*/
