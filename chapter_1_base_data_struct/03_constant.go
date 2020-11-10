package main

import (
	"fmt"
)

const PI = 3.1415926

func main() {
	fmt.Println("Hello! Go World! this is constant")
	fmt.Println(PI)
	// 常量可以省略类型
	const A string = "AAA"
	const B = "BBB"
	fmt.Println(A)
	fmt.Println(B)
	// 常量的值必须是能够在编译时就能够确定的
	const MYNAME = "Tacks" + "is so cool"
	// 内置函数 在编译期间就能获得
	const MYNAMELEN int = len(MYNAME)
	fmt.Println(MYNAME, MYNAMELEN)

	// 并行赋值
	const Mon, Tue, Wed = 1, 2, 3
	fmt.Println(Mon, Tue, Wed)

	// 枚举
	const (
		UnKnown = 0
		Female  = 1
		Male    = 2
	)
	fmt.Println(UnKnown, Female, Male)

	// 借助 iota
	const (
		// 每当 iota 在新的一行被使用时，它的值都会自动加 1
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

}

/*
常量
	1、使用const关键字定义常量，常量一旦定义将永远不会发生改变
	2、可以定义为常量的数据类型，只能是基本数据类型，例如 bool、int、float、string
	3、定义格式 const identiferName [type] = value,例如 const PI = 3.1415926
	4、可以省略类型说明符 [type]，因为编译器可以根据常量的值来推断其类型
	5、常量的值必须是能够在编译时就能够确定的，内置函数计算的值可以在编译的时候就获得
	6、常量const，可以并行赋值，可以因式分解
	7、可以用作枚举
	8、可以借助`iota`关键字,来进行枚举赋值，新的一行被使用时，它的值都会自动加 1。
*/
