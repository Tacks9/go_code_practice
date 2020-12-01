package main

import (
	"fmt"
)

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	// 匿名字段
	int // anonymous field
	// 匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体
	innerS //anonymous field
}

func main() {
	fmt.Println("This is anonymous field -------------------------匿名字段")
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	// 匿名字段int
	outer.int = 60
	// 继承innerS结构体的字段
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)

}

/*
1、结构体可以包含一个或多个 匿名（或内嵌）字段，字段的类型是必须的，此时类型就是字段的名字
2、在一个结构体中对于每一种数据类型只能有一个匿名字段
3、匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体
4、内层结构体被简单的插入或者内嵌进外层结构体。这个简单的 “继承” 机制提供了一种方式
5、外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式
6、
*/
