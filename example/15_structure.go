/*
 * @Descripttion: struct 结构体 (带类型的字段 集合)
 * @Author: tacks321@qq.com
 * @Date: 2021-02-09 17:25:51
 * @LastEditTime: 2021-02-09 17:54:40
 */

package main

import "fmt"

// 定义一个人的结构体
type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("===============struct======================")

	// 结构体类型的变量
	fmt.Println(Person{name: "Tacks", age: 18})

	// 可以省略变量名，但是必须要按照变量的顺序写
	fmt.Println(Person{"Tacks", 18})

	// 通过指定具体的变量名 可以省略一些变量，不进行赋值, 没有赋值的默认为对应类型的零值
	fmt.Println(Person{name: "Tacks"})

	// 结构体指针
	fmt.Println(&Person{name: "Tacks"})

	// 通过 .字段名 来进行获取结构体字段的值
	structPerson := Person{name: "blue", age: 20}
	fmt.Println(structPerson.name)

	// 结构体的值是可 以改变的
	structPerson.age = 99
	fmt.Println(structPerson.age)

	// 结构体和数组一样可以用new进行声明
	ms := new(Person)
	ms.name = "blue"
	ms.age = 24
	fmt.Printf("[ms] %v", ms)
}
