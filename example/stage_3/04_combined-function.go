/*
 * @Descripttion: 组合函数 combined-function
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 10:10:25
 * @LastEditTime: 2021-02-20 10:31:04
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=============================【 组合函数 combination 】======================================")

	var strs = []string{"peach", "APPLE", "Banana", "pear"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))

	// 是否有前缀是p
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	// 是否全部前缀都是p
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	// 获取所有字符串包含e的切片
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// 对所有字符串切片进行处理 得到大写
	fmt.Println(Map(strs, strings.ToUpper))
}

// 获取目标字符串中，第一个t的索引位置
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 目标字符串中是否存在 t
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 如果这些切片中的字符有一个满足 f函数 就返回 true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 如果这些切片中的字符都满足 f 那么就返回 true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		// 如果有一个不满足的就返回false
		if !f(v) {
			return false
		}
	}
	return true
}

// 返回一个切片，所有都满足条件 f 的字符串 新切片
func Filter(vs []string, f func(string) bool) []string {
	// 创建一个切片
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			// 追加切片
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 返回一个新切片，将原始切片字符串都经过 f 处理
func Map(vs []string, f func(string) string) []string {
	// 创建一个同等大小的切片
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
