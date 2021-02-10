/*
 * @Descripttion: pointer 指针
 * @Author: tacks321@qq.com
 * @Date: 2021-02-09 16:54:37
 * @LastEditTime: 2021-02-09 17:24:01
 */

package main

import "fmt"

// 传值， 参数int
func zeroval(ival int) {
	// 形参的拷贝
	ival = 0
}

// 传指针, 参数int指针
func zeroptr(iptr *int) {
	// 	*iptr 从它内存地址得到这个地址对应的当前值
	*iptr = 0
}

func main() {

	fmt.Println("===================pointer指针========================")

	i := 1
	fmt.Println("initial:", i)
	fmt.Println("pointer:", &i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过 `&i` 语法来取得 `i` 的内存地址，即指向 `i` 的指针。
	// 变量的内存地址的 引用
	zeroptr(&i) // 这里改变了i的 值
	fmt.Println("zeroptr:", i)

	// 指针也是可以被打印的。获的指向的地址，地址没有改变
	fmt.Println("pointer:", &i)

}
