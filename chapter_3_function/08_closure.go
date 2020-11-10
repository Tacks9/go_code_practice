package main

import (
	"fmt"
)

func main() {

	fmt.Println("闭包------------------------------------closure ")

	// 1 到 1 百万整数的总和的匿名函数
	// 匿名函数没有名称
	func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
		fmt.Printf("%d", sum)
		//最后的一对括号表示对该匿名函数的调用
	}()

	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

/*
1、关键字 defer （详见第 6.4 节）经常配合匿名函数使用，它可以用于改变函数的命名返回值

*/
