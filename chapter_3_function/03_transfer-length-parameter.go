package main

import (
	"fmt"
)

func main() {

	fmt.Println("传递变长参数-------------------------------------transfer-length-parameter")

	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)

	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

// 获取最小的元素
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0] // 假设第一个最小
	// 遍历整个切片 slice
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

/*

1、处理一个变长的参数，最后一个参数是采用 ...type 的形式
*/
