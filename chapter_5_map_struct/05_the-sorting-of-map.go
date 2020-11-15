package main

import (
	"fmt"
	"sort"
)

// 声明全局 map
var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("the-sorting-of-map -------------------------map 的排序")

	// map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
	// 如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序

	fmt.Println("Unsorted: ")
	for k, v := range barVal {
		fmt.Printf("key：%v, Value：%v\n", k, v)
	}

	// 声明一个切片
	keys := make([]string, len(barVal))
	i := 0
	// for循环 map，并取出来赋值给 keys切片
	for k, _ := range barVal {
		keys[i] = k
		i++
	}

	// 对切片排序
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("Sorted: ")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v \n", k, barVal[k])
	}
}
