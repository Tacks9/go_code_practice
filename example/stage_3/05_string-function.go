/*
 * @Descripttion: 字符串相关函数 标准库 strings
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 10:31:52
 * @LastEditTime: 2021-02-20 10:57:37
 */

package main

import (
	"fmt"
	// 整个别名
	s "strings"
)

// 定义一个简短的名字
var p = fmt.Println

func main() {
	fmt.Println("=============================【 字符串相关函数 string-func 】======================================")

	p("Contains  : ", s.Contains("test", "es"))
	// 统计子字符
	p("Count     : ", s.Count("test", "t"))
	// 前后缀判断
	p("HasPrefix : ", s.HasPrefix("test", "te"))
	p("HasSuffix : ", s.HasSuffix("test", "st"))
	p("Index     : ", s.Index("test", "t"))
	// 切片 转化为字符串
	p("Join      : ", s.Join([]string{"a", "b", "c", "d"}, "-"))
	// 分割字符串为 切片
	p("Split     : ", s.Split("a-b-c-d", "-"))

	p("Repeat    : ", s.Repeat("a", 5))
	p("Replace   : ", s.Replace("foo", "o", "0", -1))
	p("Replace   : ", s.Replace("foo", "o", "0", 1))

	// 大小写转化
	p("ToLower   : ", s.ToLower("TEST"))
	p("ToUpper   : ", s.ToUpper("test"))
	p("...")
}
