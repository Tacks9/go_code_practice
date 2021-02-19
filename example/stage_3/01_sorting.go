/*
 * @Descripttion: sorting 排序 (内置数据类型，自定义数据类型)
 * @Author: tacks321@qq.com
 * @Date: 2021-02-19 09:57:50
 * @LastEditTime: 2021-02-19 18:19:47
 */

package main

import (
	"fmt"
	"sort"
)

// 自定义数据类型的排序规则

// 定义别名，用来表示字符串数组
type ByLength []string

// 自定义排序sort, 我们需要实现 sort.Interface 中的 三个方法 Len() Swap() Less(), 这样就可以用sort排序
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	// 比较两个长度，按照字符串的长度来进行排序
	return len(s[i]) < len(s[j])
}

func main() {
	fmt.Println("=============================【 排序 sorting 】======================================")

	// 内置数据结构排序

	fmt.Println("[字符串数组]")
	strs := []string{"d", "a", "c", "e", "A"}
	fmt.Println("strs Original:", strs)
	// 排序 原地操作
	sort.Strings(strs)
	fmt.Println("strs Sorted  :", strs)

	fmt.Println("[整型数组]")
	ints := []int{2, 0, 11, 99, 9, 8, 10}
	fmt.Println("ints Original :", ints)
	sort.Ints(ints)
	fmt.Println("ints Sorted   :", ints)

	fmt.Println("[判断是否有序]")
	isSort := sort.IntsAreSorted(ints)
	fmt.Println("ints is Sort :", isSort)

	fruits := []string{"peach", "banana", "kiwi"}
	fmt.Println("fruits Original:", fruits)
	// 先强制转化成我们定义的类型
	// 然后采用sort.Sort进行排序
	sort.Sort(ByLength(fruits))
	fmt.Println("fruits Sorted  :", fruits)

}
