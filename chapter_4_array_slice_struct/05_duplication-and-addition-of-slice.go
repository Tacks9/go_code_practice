package main

import (
	"fmt"
)

func main() {

	fmt.Println("切片的复制与追加-------------------------------------duplication-and-addition-of-slice")

	sl_from := []int{1, 2, 3} // 一开始切片长度 3
	sl_to := make([]int, 10)  // 想要扩展到 10
	n := copy(sl_to, sl_from) // 首先把from的切片 复制到to  // 函数返回值 成功复制的元素个数
	/*
		func copy(dst, src []T)
		类型为 T 的切片从源地址 src 拷贝到目标地址 dst
		覆盖 dst 的相关元素，并且返回拷贝的元素个数
	*/
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3 //

	sl3 := []int{1, 2, 3} // 原来3个长度切片
	//  0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片 （追加的元素必须和原切片的元素同类型）
	sl3 = append(sl3, 4, 5, 6) // 追加 多个元素
	// append 方法总是返回成功
	// 如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。
	// 因此，返回的切片可能已经指向一个不同的相关数组了
	fmt.Println(sl3)

	fmt.Println(AppendInt(sl3, 99, 88, 77, 66)) // [1 2 3 4 5 6 99 88 77 66]

}

// 追加切片
func AppendInt(slice []int, data ...int) []int {
	m := len(slice)
	// 第二个参数扩展成一个列表
	// 追加后的大小长度
	n := m + len(data)
	// 如果大于当前切片容量 那么就需要重新申请切片
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]int, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	// 如果追加后长度 还是在容量内
	slice = slice[0:n]
	// 就在后面继续追加
	copy(slice[m:n], data)
	return slice
}

/*
一旦切片的容量已经到达上线，需要扩展容量，就需要用到切片的复制跟追加

1、如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来
2、拷贝切片的 copy 函数、向切片追加新元素的 append 函数

*/
