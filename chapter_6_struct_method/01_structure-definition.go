// 2020/11/16 周一

package main

import (
	"fmt"
	"strings"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToLower(p.lastName)
}

func main() {
	fmt.Println("This is structure-definition -------------------------结构体定义")
	// 使用 new 函数给一个新的结构体变量分配内存
	ms := new(struct1)
	// 使用点号符可以获取结构体字段的值：structname.fieldname
	ms.i1 = 9
	ms.f1 = 18.0
	ms.str = "Tacks"
	// 可以使用点号符给字段赋值：structname.fieldname = value
	fmt.Printf("ms.i1 int : %d\n", ms.i1)
	fmt.Printf("ms.f1 float32 : %f\n", ms.f1)
	fmt.Printf("ms.str string : %s\n", ms.str)
	// 类似使用 %v 选项
	fmt.Printf("ms struct : %v\n", ms)
	// 打印一个结构体的默认输出可以很好的显示它的内容
	fmt.Println(ms)

	fmt.Println("------------------------------------------------------ 结构体指针 ")

	type myStruct struct {
		i int
	}
	// v 是结构体类型变量
	var v myStruct
	// p 是指向一个结构体类型变量的指针
	var p *myStruct //  初始化指针的时候会为指针 p 的值赋为 nil

	fmt.Println(v) // {0}
	fmt.Println(p) // <nil>
	// v.i = 1
	// p.i = 1 // 在这里进行赋值，go会报错  invalid memory address or nil pointer dereference
	// 因为p代表的是 *p 的地址
	// nil 的话系统还并没有给 *p 分配地址, 给p 赋值肯定会出错

	// [解决办法]： 指针赋值前可以先创建一块内存分配给赋值对象
	p = new(myStruct)

	v.i = 1
	p.i = 2         // 这样就不会报错了
	fmt.Println(v)  // {1}
	fmt.Println(p)  // &{2}
	fmt.Println(&p) // p的地址 0xc000006030

	fmt.Println("------------------------------------------------------ 初始化结构体 ")

	// 指向它的指针
	// 初始化一个结构体实例 &struct1{a, b, c} 是一种简写，底层仍然会调用 new ()，这里值的顺序必须按照字段顺序来写
	// new(Type) 和 &Type{} 是等价的
	sa := &struct1{100, 99.99, "good"}

	// 结构体类型实例
	var sb struct1
	sb = struct1{60, 59.9, "bad"}
	fmt.Println(sa)
	fmt.Println(sb)

	fmt.Println("------------------------------------------------------ 结构体字面量初始化 结构体类型实例  字段可以省略 ")

	type Interval struct {
		start int
		end   int
	}
	// 值必须以字段在结构体定义时的顺序给出
	intrA := Interval{1, 3}
	// 显示了另一种方式，字段名加一个冒号放在值的前面，这种情况下值的顺序不必一致
	intrB := Interval{end: 5, start: 1}
	// 某些字段还可以被忽略掉
	intrC := Interval{end: 5}
	fmt.Println(intrA)
	fmt.Println(intrB)
	fmt.Println(intrC)

	fmt.Println("------------------------------------------------------ 对象本身是可改变的 ")

	//[1] 结构体字面量初始化
	var per1 Person
	per1.firstName = "Zhou"
	per1.lastName = "JieLun"
	upPerson(&per1)
	fmt.Println(per1) // {ZHOU jielun}
	fmt.Printf("per1 name is %s %s\n", per1.firstName, per1.lastName)

	// new(Type) 和 &Type{} 是等价的

	//[2] new 初始化 得到一个指针
	per2 := new(Person)
	per2.firstName = "Zhou"
	per2.lastName = "JieLun"
	// 可以直接通过指针，像 pers2.lastName="JieLun" 这样给结构体字段赋值
	(*per2).lastName = "JieLun" // 合法操作
	upPerson(per2)
	fmt.Println(per2) // &{ZHOU jielun}
	fmt.Printf("per2 name is %s %s\n", per2.firstName, per2.lastName)

	//[3]
	// 直接得到的就是一个指针类型的结构体
	per3 := &Person{"Zhou", "JieLun"}
	upPerson(per3)
	fmt.Println(per3) // &{ZHOU jielun}
	fmt.Printf("per3 name is %s %s\n", per3.firstName, per3.lastName)

	fmt.Println("------------------------------------------------------ 类型转换遵循严格的规则 ")

	type number struct {
		f float32
	}

	type nb number
	// 结构体字面量初始化
	a := number{10.0}
	b := nb{10.0}
	var c = number(b)
	fmt.Println(a, b, c)
	// 指针类型的结构体
	d := &number{100.0}
	e := new(number)
	e.f = 100.0
	fmt.Println(d, e)

}

/*
[结构体]
type identifier struct {
    field1 type1
    field2 type2
}

1、type T struct {a, b int} 简单的结构体
2、数组可以看作是一种结构体类型，不过它使用下标而不是具名的字段
3、使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针 var t *T = new(T)
4、Println 打印一个结构体的默认输出可以很好的显示它的内容，类似使用 %v 选项
5、选择器（selector）。 可以使用点号符给字段赋值：structname.fieldname = value 。使用点号符可以获取结构体字段的值：structname.fieldname
6、new(Type) 和 &Type{} 是等价的
7、结构体和它所包含的数据在内存中是以连续块的形式存在的
8、结构体转换，类型转换遵循严格的规则，当为结构体定义了一个 alias 类型时，此结构体类型和它的 alias 类型都有相同的底层类型
*/
