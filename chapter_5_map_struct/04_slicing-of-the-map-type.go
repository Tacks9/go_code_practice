package main

import "fmt"

func main() {
	fmt.Println("slicing-of-the-map-type -------------------------map 类型的切片")

	// 获取一个 map 类型的切片
	items := make([]map[int]int, 5)
	for i := range items {
		// 通过索引使用切片的 map 元素
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
	// Version A: Value of items: [map[1:2] map[1:2] map[1:2] map[1:2] map[1:2]]

}
