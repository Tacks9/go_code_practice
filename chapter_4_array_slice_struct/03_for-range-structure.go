package main

import (
	"fmt"
)

func main() {

	fmt.Println("-------------------------------------for range")
	// 创建一个大小为4的int切片，同时创建好相关数组
	var slice1 []int = make([]int, 4)

	// 切片赋值
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	// 遍历切片
	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}

	fmt.Println("-------------------------------------for range")
	// 字符串切片
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	var season string
	// _ 忽略索引
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

}

/*

遍历切片

for ix, value := range slice1 {
    ...
}

1、第一个返回值 ix 是数组或者切片的索引，
2、第二个是在该索引位置的值；他们都是仅在 for 循环内部可见的局部变量。
3、value 只是 slice1 某个索引位置的值的一个拷贝，不能用来修改 slice1 该索引位置的值。


*/
