package main

import (
	"fmt"
	"io"
	"log"
)

func main() {

	fmt.Println("defer 和追踪-------------------------------------defer and tracking")

	fmt.Println("defer ")

	function1()

	fmt.Println("多个defer")

	f() // 4 3 2 1 0

	fmt.Println("\ndefer追踪")
	b()

	func1("GOGOGO") // 2020/11/10 14:04:08 func("GOGOGO") = 9, EOF
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2() // 在return之前执行该语句
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function!")
}

func f() {
	for i := 0; i < 5; i++ {
		// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
		defer fmt.Printf("%d ", i)
	}
}

// 追踪
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

// 离开
func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func(%q) = %d, %v", s, n, err)
	}()
	return 9, io.EOF
}

/*

1、关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
2、当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
3、可以做函数的收尾工作，关闭文件流、解锁一个加锁的资源、打印最终报告、关闭数据库链接
4、使用 defer 语句实现代码追踪
5、log日志包，io相关包
*/
