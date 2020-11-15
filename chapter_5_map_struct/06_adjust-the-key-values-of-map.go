package main

import (
	"fmt"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("adjust-the-key-values-of-map ------------------------- 将 map 的键值对调")
	// 这里对调是指调换 key 和 value。如果 map 的值类型可以作为 key 且所有的 value 是唯一的

	// 申请一个切片
	invMap := make(map[int]string, len(barVal))

	// 键值互换
	for k, v := range barVal {
		invMap[v] = k
	}

	fmt.Println("inverted:")

	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / \n", k, v)
	}
}
