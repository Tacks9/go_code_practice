package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {

	fmt.Println("使用闭包调试------------------------------------use-closure-debugging ")

	// whereLine() 闭包函数来打印函数执行的位置
	whereLine := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	whereLine()

}

/*
1、使用 runtime 或 log 包中的特殊函数来分析和调试复杂的程序
*/
