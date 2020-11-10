package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("test the error of multiple return function")

	var orig string = "ABC"
	var newS string
	// 当前64位处理器
	fmt.Printf("The size of ints is : %d \n", strconv.IntSize)

	// 字符串转成整型
	an, err := strconv.Atoi(orig)
	// comma,ok 模式
	if err != nil {
		// 转化失败
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		// os.Exit(1) // 退出代码 1 可以使用外部脚本获取到
		// return
	}

	fmt.Printf("The integer is %d\n", an)
	// 整型转成字符串
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	ount, err := fmt.Println(orig)
	fmt.Printf("Println is: %s err is %d\n", ount, err)

}

/*


1、Go 语言的函数经常使用两个返回值来表示执行是否成功：返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败
2、Go 语言中的错误类型为 error: var err error
3、comma,ok 模式
4、

*/
