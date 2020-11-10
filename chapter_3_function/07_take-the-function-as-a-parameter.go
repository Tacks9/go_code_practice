package main

import (
	"fmt"
)

func main() {

	fmt.Println("将函数作为参数-------------------------------------take-the-function-as-a-parameter")

	callback(1, add)
}

func add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

// 回调
func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes add(1, 2)
}

/*
1、函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调。

*/
