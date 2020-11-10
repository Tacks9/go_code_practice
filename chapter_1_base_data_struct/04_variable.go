package main

import (
	"fmt"
	"os"
	"runtime"

	// 导入了包 fourinclude(init.go 目录为./fourinclude/init.go）

	"./fourinclude"
)

// 并且使用包的变量 Pi
var twoPi = 2 * fourinclude.Pi

func main() {
	fmt.Println("Hello! Go World! this is variable! ")

	// 显示声明变量类型
	var goOs string = runtime.GOOS
	fmt.Println("The Operating system is: %s\n", goOs)
	// go在运行时候自动识别path为字符串类型
	path := os.Getenv("PATH")
	fmt.Println("Path is %s\n", path)

	name := "tacks"
	fmt.Println(name)

	// 计算
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586

}

/*

变量
	1、声明，变量的类型放在变量的名称之后，`var a int = 1`
	2、使用var声明变量，因式分解关键字的写法一般用于声明全局变量
	3、当一个变量被声明之后，系统自动赋予它该类型的零值，int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil
	4、变量标识符，也就是命名规则，遵循骆驼命名法
	5、全局变量希望能够被外部包所使用，标识符开头大写，否则小写
	6、变量的声明与赋值，var identifier [type] = value
	7、某些情况下， Go 是可自动推断类型
值类型与引用类型
	1、int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
	2、数组 array 、结构 struct 这些复合类型也是值类型，值类型的变量的值存储在栈中
	3、&i 来获取变量 i 的内存地址
	4、指针、slice、maps、channel属于引用类型，被引用的变量会存储在堆中
	5、值类型，使用等号 i = j 将一个变量的值赋值给另一个变量时，实际上是将 i 的值进行了拷贝
	6、引用类型，当使用赋值语句 r2 = r1 时，只有引用（地址）被复制, 如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容
打印
	1、fmt.Printf()  打印格式化字符串
	2、fmt.Sprintf() 返回格式化后的字符串
	3、fmt.Print()   打印多个字符串 中间会空格
	4、fmt.Println() 打印多个字符串 中间会空格 结尾会换行
赋值
	1、简写：，a := 50，变量类型将由编译器自动推断
	2、并行赋值，多个变量，a, b, c = 5, 7, "abc"
	3、交换两个变量的值 a, b = b, a ，非常方便，不需要使用交换函数
	4、空白标识符 _ 也被用于抛弃值，值 5 在：_, b = 5, 7 中被抛弃
	5、一个函数返回多个返回值时，val, err = Func1(var1)
init函数
	1、init 函数中初始化，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高
	2、每个源文件都只能包含一个 init 函数。初始化总是以单线程执行，并且按照包的依赖关系顺序执行

*/
