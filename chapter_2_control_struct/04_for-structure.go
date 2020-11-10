package main

import (
	"fmt"
)

func main() {
	fmt.Println("for structure")

	//---------------------------------------------------------- 【1】 基于计数器的迭代
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	// 打印字符串中每个字符
	str := "Go is a beautiful language!" // ASCII 编码的字符占用 1 个字节
	strLen := len(str)
	fmt.Printf("The length of str is: %d\n")
	for ix := 0; ix < strLen; ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	/*
	   	// 1: 循环 15 次并进行打印
	   	for i := 0; i < 15; i++ {
	   		fmt.Printf("The counter is at %d\n", i)
	   	}
	   	// 2:
	   	i := 0
	   START:
	   	fmt.Printf("The counter is at %d\n", i)
	   	i++
	   	if i < 15 {
	   		// goto 语句重写循环
	   		goto START
	   	}
	*/
	/*
		//---写一个从 1 打印到 100 的程序
		// 每当遇到 3 的倍数时，不打印相应的数字 打印一次 "Fizz"
		// 遇到 5 的倍数时，打印 Buzz 而不是相应的数字
		// 于同时为 3 和 5 的倍数的数，打印 FizzBuzz
		const (
			FIZZ     = 3
			BUZZ     = 5
			FIZZBUZZ = 15
		)
		for i := 0; i <= 100; i++ {
			switch {
			case i%FIZZBUZZ == 0:
				fmt.Println("FizzBuzz")
			case i%FIZZ == 0:
				fmt.Println("Fizz")
			case i%BUZZ == 0:
				fmt.Println("Buzz")
			default:
				fmt.Println(i)
			}
		}
	*/
	//---- 打印三角形
	line := 10
	for i := 1; i < line; i++ {
		// 先打印空格
		for k := 1; k <= line-i; k++ {
			fmt.Print(" ")
		}
		// 打印*
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//----------------------------------------------------------【2】 基于条件判断的迭代
	// 没有头部的条件判断迭代，有点类似其他语言的while条件
	// 没有初始化语句和修饰语句的 for 结构 ，可以省略 ;;
	i := 10
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}
	//----------------------------------------------------------【3】 无限循环
	for {
		i++
		if i == 10 {
			fmt.Printf("The variable i is 10: %d\n", i)
			break
		}
	}
	//----------------------------------------------------------【4】for range
	fmt.Println()
	str2 := "中国人"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		// char  始终为 对应索引的值拷贝，因此它一般只具有只读性质
		// 自动根据 UTF-8 规则识别 Unicode 编码的字符，一个中文3个字节
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		// 一个字符串是 Unicode 编码的字符
		fmt.Printf("%-2d   %d   %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}

	// ----------------------------------------------------------- 练习
	// 练习1
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v) // 0 0 0 0 0
		v = 5
	}
	// 练习2 死循环
	for i := 0; ; i++ {
		// fmt.Println("Value of i is now:", i)
	}
	// 练习3 死循环
	for i := 0; i < 3; {
		// fmt.Println("Value of i:", i)
	}

}

/*
for
	1、基于计数器的迭代 for 初始化语句; 条件语句; 修饰语句 {}
	2、for后面不需要用()括号 ，左花括号 { 必须和 for 语句在同一行
	3、
*/
