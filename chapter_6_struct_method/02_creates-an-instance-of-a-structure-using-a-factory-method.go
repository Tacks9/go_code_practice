package main

import (
	"fmt"
	"unsafe"

	"./matrix"
)

/*
结构体
	如果 File 是一个结构体类型，那么表达式 new(File) 和 &File{} 是等价的
*/
type File struct {
	fd   int    // 文件描述符
	name string // 文件名称
}

/*
结构体工厂
	工厂实例化了类型的一个对象
*/
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main() {
	fmt.Println("This is creates-an-instance-of-a-structure-using-a-factory-methodn -------------------------使用工厂方法创建结构体实例")

	fmt.Println("---------------------------------------------------------结构体类型对应的工厂方法 New")

	// 结构体类型 T 的一个实例占用了多少内存  24
	fmt.Println(unsafe.Sizeof(File{}))
	f := NewFile(10, "./test.txt") //调用
	fmt.Println(f)

	fmt.Println("---------------------------------------------------------如何强制使用工厂方法")

	// wrong := new(matrix.matrix)     // 编译失败（matrix 是私有的）
	right := matrix.NewMatrix("Factory Matrix") // 实例化 matrix 的唯一方式
	fmt.Println(right)

	fmt.Println("---------------------------------------------------------new  make区别")
	type Foo map[string]string
	type Bar struct {
		thingOne string
		thingTwo int
	}

	// new结构体 初始化
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1
	fmt.Println(y) // &{hello 1}

	// NOT OK 结构体不能使用make
	// z := make(Bar) // 编译错误：cannot make type Bar
	// (*z).thingOne = "hello"
	// (*z).thingTwo = 1

	// OK map哈希 可以make
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"
	fmt.Println(x) // map[x:goodbye y:world]

	// Not Ok map哈希 在new的时候是nil初始化， 使用数据填充它，将会引发运行时错误
	// 因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内
	// u := new(Foo)
	// (*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	// (*u)["y"] = "world"
}

/*

1、Go 中实现 “构造子工厂” 方法，为了方便通常会为类型定义一个工厂，按惯例，工厂的名字以 new 或 New 开头
2、如果 File 是一个结构体类型，那么表达式 new(File) 和 &File{} 是等价的
3、工厂实例化了类型的一个对象，类似其他面向对象语言的中 基于类的 OOP
4、结构体类型 T 的一个实例占用了多少内存 ，size := unsafe.Sizeof(T{}) ，需要引入 unsafe包
5、强制使用工厂方法，禁止使用 new 函数，强制用户使用工厂方法，从而使类型变成私有的

*/
