package main

import (
	"fmt"
	"time"
)

// 求解第41的斐波那契
const LIM = 41

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	/// 时间计算
	start := time.Now()

	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	// 将第 n 个数的值存在数组中索引为 n 的位置
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
