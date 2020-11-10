package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("应用闭包：将函数作为返回值------------------------------------application-closure-function-as-a-return-value ")

	p2 := add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))

	TwoAdder := adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	// 封装工厂函数
	addBmp := makeAddSuffix(".bmp")
	addJpeg := makeAddSuffix(".jpeg")

	fmt.Printf("addBmp file is: %s\n", addBmp("file"))   // returns: file.bmp
	fmt.Printf("addJpeg file is: %s\n", addJpeg("file")) // returns: file.jpeg

}

func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

// 一个返回值为另一个函数的函数可以被称之为工厂函数
func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

/*
1、可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，
2、闭包在 Go 语言中非常常见，常用于 goroutine 和管道操作
3、函数也是一种值，因此很显然 Go 语言具有一些函数式语言的特性
*/
