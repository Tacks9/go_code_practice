package main

import (
	"fmt"
)

func main() {

	fmt.Println("递归函数-------------------------------------recursive-function")

	for i := 0; i <= 10; i++ {
		res := fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, res)
	}

	deNum(10)
}

/*
斐波那契
*/
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}

// 从10打印到1
func deNum(n int) {
	if n >= 1 {
		deNum(n - 1)
	}
	fmt.Printf("%d ", n)
}

/*
1、当一个函数在其函数体内调用自身，则称之为递归

*/
