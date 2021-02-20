/*
 * @Descripttion: 正则表达式 regular-expression
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 11:30:44
 * @LastEditTime: 2021-02-20 11:52:39
 */

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("=============================【 正则表达式 regular 】======================================")

	// 判断一个字符串是否符合正则
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 优化一个Compile的Regexp
	r, _ := regexp.Compile("p([a-z]+)ch")
	// 等同第一个 表达式
	fmt.Println(r.MatchString("peach"))

	// 查找字符串
	fmt.Println(r.FindString("peach punch")) // peach

	// 返回第一次匹配到字符串的索引位置 [开始 结束]
	fmt.Println((r.FindStringIndex(" peach punch"))) // [1 6]

	// 返回第一次 完全匹配 局部匹配的索引位置
	fmt.Println(r.FindStringSubmatchIndex(" peach punch")) // [1 6 2 4]

	// 返回所有的匹配项
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	//  `All` 同样可以对应到上面的 所有项 完全匹配 局部匹配
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// 返回匹配的内容 限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	// 上面的例子中，我们使用了字符串作为参数，并使用了如 `MatchString` 这样的方法。
	// 我们也可以提供 `[]byte` 参数并将 `String` 从函数命中去掉。
	fmt.Println(r.Match([]byte("peach")))

	// 创建正则表达式常量时，可以使用 `Compile` 的变体 `MustCompile` 。
	// 因为 `Compile` 返回两个值，不能用于常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// `regexp` 包也可以用来替换部分字符串为其他值。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// `Func` 变量允许传递匹配内容到一个给定的函数中，
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
