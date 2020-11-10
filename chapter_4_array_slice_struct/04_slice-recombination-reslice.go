package main

import (
	"fmt"
)

func main() {

	fmt.Println("-------------------------------------slice-recombination-reslice")

	// 初始化切片长度为 0
	slice1 := make([]int, 0, 10)
	// 相关数组长度 10
	for i := 0; i < cap(slice1); i++ {
		// 每次扩展1个长度
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

}

/*

make创建一个切片

slice1 := make([]type, start_length, capacity)

start_length 作为切片初始长度而 capacity 作为相关数组的长度

切片在达到容量上限后可以扩容。改变切片长度的过程称之为切片重组 reslicing

做法如下：slice1 = slice1[0:end]，其中 end 是新的末尾索引（即长度）。
*/
