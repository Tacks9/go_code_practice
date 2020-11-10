package main

import (
	"fmt"
)

func main() {
	fmt.Println("函数参数与返回值-------------------------------------function-parameters-and-return-values")

	fmt.Println("--------------------------------------------------非命名返回值 与 命名返回值")

	// 非命名返回值 与 命名返回值
	fmt.Println(getX2AndX3(1))
	fmt.Println(getX2AndX3_2(1))

	// 空白符用来匹配一些不需要的值
	a, _ := getX2AndX3_2(1)
	fmt.Println(a)

	fmt.Println("--------------------------------------------------命名返回值，返回多个参数")
	var min, max int
	min, max = minMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)

	fmt.Println("--------------------------------------------------改变外部变量")
	n := 0
	// reply 是一个指向 int 变量的指针
	reply := &n
	multiply(10, 5, reply)           // 我们在函数内修改了这个 reply 变量的数值
	fmt.Println("Multiply:", *reply) // Multiply: 50

}

// ------------------------------------- 非命名返回值 (int, int)
// 当需要返回多个非命名返回值时，需要使用 () 把它们括起来
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// ------------------------------------- 命名返回值 (x2 int, x3 int)
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// 只需要一条简单的不带参数的 return 语句
	return
}

// ------------------------------------ 返回最小值最大值 minmax
func minMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}

// ------ 传递指针
func multiply(a, b int, reply *int) {
	// reply 是一个指向 int 变量的指针
	*reply = a * b
}

/*
1、多值返回是 Go 的一大特性
2、niladic 函数，定义没有形参名的函数，只有相应的形参类型 func f(int, int, float64),例如main.main()

按值传递（call by value） 按引用传递（call by reference）
	1、函数参数传值类型，按值传递、按引用传递
	2、Go 默认使用按值传递来传递参数，使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量 （ 例如 Function(arg1)）
	3、引用传递，希望函数可以直接修改参数的值，则需要将参数的地址传递给函数（ 例如 Function(&arg1)）
	4、像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递
	5、几乎在任何情况下，传递指针（一个 32 位或者 64 位的值）的消耗都比传递副本来得少

命名的返回值
	1、当需要返回多个非命名返回值时，需要使用 () 把它们括起来
	2、命名返回值作为结果形参（result parameters）被初始化为相应类型的零值，当需要返回的时候，我们只需要一条简单的不带参数的 return 语句 ，需要使用 () 把它们括起来

空白符
	1、空白符用来匹配一些不需要的值，然后丢弃掉

改变外部变量
	1、函数参数，是指针类型，这样不需要对形参进行拷贝，可以节省内存
	2、传递指针，可以直接修改形参的内容，并且函数返回的时候不需要使用return
	3、十分小心那些可以改变外部变量的函数

*/
