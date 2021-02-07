/*
 * @Descripttion: 循环结构 for
 * @Author: tacks321@qq.com
 * @Date: 2021-02-07 14:06:46
 * @LastEditTime: 2021-02-07 16:58:39
 */

package main

import "fmt"

func main() {
	fmt.Println("========================【for循环结构】=============================")

	// 类似 while循环
	i := 1
	for i <= 3 {
		fmt.Print(i, " ")
		i = i + 1
	}

	// 经典for循环
	for j := 4; j <= 6; j++ {
		fmt.Print(j, " ")
	}

	// 借助break的死循环
	for {
		fmt.Println("LOOP!")
		break
	}

	fmt.Println(" ")
	// 使用continue
	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n)
	}

	// 字符串类型打印

	// 打印 英文
	str := "Go is a beautiful language!" // ASCII 编码的字符占用 1 个字节
	strLen := len(str)
	fmt.Printf("The length of str is: %d\n")
	for ix := 0; ix < strLen; ix++ {
		fmt.Printf("%c ", str[ix])
	}

	// 打印 中文
	fmt.Println()
	str2 := "中国人"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		// char  始终为 对应索引的值拷贝，因此它一般只具有只读性质
		// 自动根据 UTF-8 规则识别 Unicode 编码的字符，一个中文3个字节
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

}
