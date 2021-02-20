/*
 * @Descripttion: panic 异常处理
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 09:32:09
 * @LastEditTime: 2021-02-20 09:42:22
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=============================【 异常处理 panic 】======================================")

	// 预定义一个错误
	panic("a problem!")

	// 创建一个文件
	_, err := os.Create("/tmp/file")

	// GO 习惯通过 返回值来标示错误
	if err != nil {
		// 程序会引起 一个输出错误/GO运行时栈信息/非零的状态码
		panic(err)
	}

}
