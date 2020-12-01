/*
 * @Description:
 * @Author: Tacks
 * @Date: 2020-11-17 09:35:45
 */
// 2020-11-17 周二

package main

import (
	"fmt"
)

// 结构体
type TwoInts struct {
	a int
	b int
}

// 切片类型
type IntVector []int

func main() {
	fmt.Println("This is method -------------------------方法")

	fmt.Println("--------------------------------------------------------------结构体类型 方法")

	// 结构体实例,	// 使用 new 函数给一个新的结构体变量分配内存
	two1 := new(TwoInts)
	two1.a = 10
	two1.b = 20

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	// 结构体字面量初始化 值必须以字段在结构体定义时的顺序给出
	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())

	fmt.Println("--------------------------------------------------------------非结构体类型 方法")
	fmt.Println(IntVector{1, 2, 3}.Sum()) // 输出是6

}

// 结构体类型 求和方法

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

// 切片类型 上方法
// 函数注释

/**
 * @name:
 * @description: xxx
 * @param {IntVector} v
 * @return {*}
 */
func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	// 命名返回值
	// 只需要一条简单的不带参数的 return 语句
	return
}

/*
方法格式
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

1、结构体就像是类的一种简化形式
2、Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。
3、接收者类型可以是（几乎）任何类型，不能是接口类型（抽象定义），不能是指针类型
4、因为方法是函数，所以同样的，不允许方法重载
5、对于一个类型只能有一个给定名称的方法。但是如果基于接收者类型，具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在
6、定义格式，在方法名之前，func 关键字之后的括号中指定 receiver
7、如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()
8、

*/
