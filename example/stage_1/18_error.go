/*
 * @Descripttion: 错误相关error
 * @Author: tacks321@qq.com
 * @Date: 2021-02-10 11:12:46
 * @LastEditTime: 2021-02-10 15:00:29
 */

package main

import (
	"errors"
	"fmt"
)

// 函数
// errors使用一个独立、明确的返回值来传递错误信息
func ff(arg int) (int, error) {
	if arg == 0 {
		// 采用 new 来构造一个使用给定错误信息的 error值
		return -1, errors.New("arg can not work with 0")
	}
	// nil 表示没有错误
	return arg, nil
}

// 自定义错误类型
type argError struct {
	arg  int
	prob string
}

// 封装一个结构体的方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func gg(arg int) (int, error) {
	if arg == 0 {
		// 这样error中可以提供两个字段
		return -1, &argError{arg, "arg can not work with 0 "}
	}
	return arg, nil
}

func main() {
	fmt.Println("========================error错误================================")

	// ff() 函数
	for _, i := range []int{1, 2, 3, 0} {
		if val, err := ff(i); err != nil {
			fmt.Println("ff() failed:", err)
		} else {
			fmt.Println("ff() work:", val)
		}
	}

	// gg() 函数
	for _, i := range []int{1, 2, 3, 0} {
		if val, err := gg(i); err != nil {
			fmt.Println("gg() failed:", err)
		} else {
			fmt.Println("gg() work:", val)
		}
	}
}
