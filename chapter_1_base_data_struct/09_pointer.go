package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello! Go World! this is pointer! ")

	fmt.Println("---------------------------------------------------pointer 指针")

	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var intP *int // 指向 int 的指针
	intP = &i1    // 此时 intP 指向 i1
	// intP  是地址
	// *intP 地址存储的值
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string

}

/*
指针

	1、程序在内存中存储它的值，每个内存块（或字）有一个地址，通常用十六进制数表示，如：0x6b0820
	2、Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	3、指针的格式化标识符为 %p
	4、当一个指针被定义后没有分配到任何变量时，它的值为 nil
*/
