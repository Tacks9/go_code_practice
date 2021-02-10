/*
 * @Descripttion: 接口 interface
 * @Author: tacks321@qq.com
 * @Date: 2021-02-10 10:28:15
 * @LastEditTime: 2021-02-10 11:08:58
 */

package main

import (
	"fmt"
	"math"
	"reflect"
)

// 接口定义
/*
1、接口一组方法签名的集合
2、接口是一种变量类型
*/

type Geometry interface {
	area() float64
	perim() float64
}

// 矩形 结构体定义
type Rect struct {
	width  float64
	height float64
}

// 圆形 结构体定义
type Circle struct {
	radius float64
}

// go中要想实现一个接口，就要实现接口中所有的方法

// 矩形 方法
func (r Rect) area() float64 {
	return r.height * r.width
}
func (r Rect) perim() float64 {
	return (r.height + r.width) * 2
}

// 非接口r方法
func (r Rect) say() string {
	return "我是矩形"
}

// 圆形 方法
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perim() float64 {
	return math.Pi * c.radius * 2
}

// 非接口方法
func (c Circle) say() string {
	return "我是圆形"
}

// 通用方法
// 传入参数是接口类型，
func measure(g Geometry) {
	// 调用对应的方法，因为接口体都实现了接口的方法

	// reflect.TypeOf() 通过反射获取当前变量的类型
	fmt.Println("[", reflect.TypeOf(g), "]: ", g, g.area())
	fmt.Println("[", reflect.TypeOf(g), "]: ", g, g.perim())

	// 这个地方就不能使用 g来调用say
	// fmt.Println("[", reflect.TypeOf(g), "]: ", g, g.say())
}
func main() {
	fmt.Println("==============接口 interface ======================")

	rs := Rect{width: 2, height: 3}
	cs := Circle{radius: 2}

	// 由于 Rect \ Circle 结构体都已经实现 measure参数接口的方法，因此可以传入进去
	// Go 中没有 implements 关键字，而是自动确定类型是否满足接口
	measure(rs)
	measure(cs)

	// interface{} 类型 空接口，表示可以传入任何类型的值
}
