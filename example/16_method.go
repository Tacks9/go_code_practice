/*
 * @Descripttion: 结构体类型中定义方法Method, 特殊的函数
 * @Author: tacks321@qq.com
 * @Date: 2021-02-10 10:00:31
 * @LastEditTime: 2021-02-10 10:19:43
 */
package main

import "fmt"

// 矩形 结构体
type Rectangle struct {
	width  int
	height int
}

// 方法：接收者（receiver）上的一个函数
// 如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()

// area() 方法  传入一个接收器 指针类型 *Rectangle
func (r *Rectangle) area() int {
	return r.height * r.width
}

// perimeter() 方法 传入一个接收器 值类型 Rectangle
func (r Rectangle) perimeter() int {
	return (r.height + r.width) * 2
}

func main() {
	fmt.Println("========================method===============================")

	rs := Rectangle{width: 10, height: 20}
	fmt.Println(rs)
	// 调用结构体方法
	fmt.Println("[area]: ", rs.area())
	fmt.Println("[perimeter]:  ", rs.perimeter())

	//  Go 自动处理方法调用时的值和指针之间的转化
	// 但是我们可以采用 指针来调用方法来避免在方法调用时产生一个拷贝
	rp := &rs
	fmt.Println(rp)
	fmt.Println("[area]: ", rp.area())
	fmt.Println("[perimeter]: ", rp.perimeter())

}
