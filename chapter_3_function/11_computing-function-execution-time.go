package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("计算函数执行时间------------------------------------computing-function-execution-time ")

	start := time.Now()
	longCalculation()

	for i := 0; i <= 40; i++ {
		res := fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, res)
	}

	end := time.Now()
	// 结束时间 - 开始时间
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}

func longCalculation() {
	for i := 1; i <= 1000000; i++ {
		// fmt.Println("The long Calculation")
	}
	fmt.Println("The long Calculation")

}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}
