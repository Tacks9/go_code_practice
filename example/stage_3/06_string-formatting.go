/*
 * @Descripttion: 字符串相关格式化输出
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 10:58:49
 * @LastEditTime: 2021-02-20 11:28:10
 */
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	fmt.Println("=============================【  字符串格式化 format 】======================================")

	fmt.Println("[Printf     os.Stdout]")
	p := point{1, 2}
	// ===============  打印结构体变量
	fmt.Printf("%v \n", p)
	// 包含结构体字段名
	fmt.Printf("%+v \n", p)
	// 值的运行源代码片段
	fmt.Printf("%#v \n", p)
	// 打印值的类型
	fmt.Printf("%T \n", p)
	// 指针
	fmt.Printf("%p \n", &p)

	// ==============
	// 布尔值
	fmt.Printf("%t \n", true)
	// 十进制格式化
	fmt.Printf("%d \n", 620)
	// 二进制输出
	fmt.Printf("%b \n", 1024)
	// acsii输出
	fmt.Printf("%c \n", 67)
	// 十六进制
	fmt.Printf("%x \n", 456)
	// 浮点数
	fmt.Printf("%f \n", 20.21)
	// 科学计数法
	fmt.Printf("%e \n", 123400000.0)
	fmt.Printf("%E \n", 123400000.0)
	// 字符串
	fmt.Printf("%s \n", "\"string\"")
	// base-16 编码 一个字节占据两个字符
	fmt.Printf("%x \n", "hex this")

	// ============================== 对齐
	// 右对齐，空白部分空格替代
	fmt.Printf("|%6d|%6d| \n", 12, 232)
	// 指定浮点型的输出宽度
	fmt.Printf("|%6.2f|%6.2f| \n", 1.2, 2.32)
	// 左对齐 加上-
	fmt.Printf("|%-6.2f|%-6.2f| \n", 1.2, 2.32)
	// 字符串右对齐
	fmt.Printf("|%6s|%6s| \n", "foo", "f")
	// 字符串左对齐
	fmt.Printf("|%-6s|%-6s| \n", "foo", "f")

	fmt.Println("[SPrintf 不输出      os.Stdout]")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// 格式化 写
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
