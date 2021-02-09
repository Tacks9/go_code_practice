/*
 * @Descripttion: function 函数 (签名、可变参数、多返回值)
 * @Author: tacks321@qq.com
 * @Date: 2021-02-09 15:44:55
 * @LastEditTime: 2021-02-09 16:32:53
 */
package main

import "fmt"

// 两数之和
// 接受两个 `int` 并且以 `int` 返回它们的和
func add(a int, b int) int {
	return a + b
}

// 三数之和
// 当多个连续的参数为同样类型时，最多可以仅声明最后一个参数类型
func addThree(a, b, c int) int {
	return a + b + c
}

//////////////////////////////////////////////////////////////////////【可变参数函数】

// N数之和
// 采用... 可以传入多个参数，是一个切片params
/*
可变参数函数
	即其参数数量是可变的 —— 0 个或多个。 （当参数数量未知 可以使用，我们的 Println 打印函数就支持）

	声明可变参数函数的方式是在其参数类型前带上省略符（三个点）前缀

*/
func adds(params ...int) int {
	// ... 可变参数 相当于在函数内部 隐式创建了一个这样的切片 params := []int{1,2}
	// 如果调用函数没有传入参数，则默认切片为 nil
	sum := 0
	for _, val := range params {
		sum = sum + val
	}
	return sum
}

//////////////////////////////////////////////////////////////////////【多返回值函数】
// 计算面积 周长
func getAreaAndPerimeter(length, width int) (int, int) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

// 主函数
func main() {
	fmt.Println("====================function函数==================")

	fmt.Println(add(1, 2))
	fmt.Println(addThree(1, 2, 3))

	fmt.Println(adds(1, 2))
	fmt.Println(adds(1, 2, 3))

	// 【可变参数函数】
	fmt.Println(adds(1, 2, 3, 4, 5, 6, 7, 8, 9))
	nums := []int{100, 200, 300, 400}
	fmt.Println(adds(nums...))

	// 【多返回值函数】
	fmt.Println(getAreaAndPerimeter(2, 2))

}
