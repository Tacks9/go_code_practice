// 2020-11-10 周二

package main

import (
	"fmt"
)

func main() {
	fmt.Println("声明和初始化---------------------------------declarations-and-initialization")

	var arr1 [5]int

	fmt.Println("------------------------------------------- for 循环处理数组")

	for i := 0; i < 5; i++ {
		arr1[i] = i * 2
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("Array at index %d is %d \n", i, arr1[i])
	}
	fmt.Println("------------------------------------------- for range ")

	for i := range arr1 {
		fmt.Printf("Array at index %d is %d \n", i, arr1[i])
	}

	fmt.Println("------------------------------------------- for range 循环处理字符串数组")
	stringArr := [...]string{"a", "b", "c", "d"} // 初始化数组
	for i := range stringArr {
		fmt.Println("Array item", i, "is", stringArr[i])
	}

	fmt.Println("------------------------------------------- new 创建")

	var arrP [3]int
	f(arrP)   // passes a copy of ar
	fp(&arrP) // passes a pointer to ar

	fmt.Println("------------------------------------------- 数组发生内存拷贝")
	var arrC [5]int

	for i := 0; i < len(arrC); i++ {
		arrC[i] = i * 2
	}

	arrV := arrC
	// 内存拷贝
	// 赋值后修改 arrV 不会对 arrC 生效
	arrV[2] = 100

	for i := 0; i < len(arrC); i++ {
		fmt.Printf("Array arrC at index %d is %d\n", i, arrC[i])
	}
	fmt.Println()
	for i := 0; i < len(arrV); i++ {
		fmt.Printf("Array arrV at index %d is %d\n", i, arrV[i])
	}

	fmt.Println("------------------------------------------- 数组初始化")
	var arrAge = [5]int{18, 20, 15, 22, 16}
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arrAge)      // 全部元素写出来
	fmt.Println(arrLazy)     // 可同样可以忽略，从技术上说它们其实变化成了切片
	fmt.Println(arrKeyValue) // 只有索引 3 和 4 被赋予实际的值，其他元素都被设置为空的字符串

	fmt.Println("------------------------------------------- 数组指针")
	for i := 0; i < 3; i++ {
		// 任意数组常量的地址来作为指向新实例的指针
		fp(&[3]int{i, i * i, i * i * i})
	}

	fmt.Println("------------------------------------------- 多维数组")

	const (
		WIDTH  = 5
		HEIGHT = 5
	)
	// 别名
	type pixel int
	var screen [WIDTH][HEIGHT]pixel
	// 初始化数组
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = pixel(x + y)
		}
	}
	fmt.Println(screen[0])
	fmt.Println(screen[1])
	fmt.Println(screen[2])
	fmt.Println(screen[3])
	fmt.Println(screen[4])

	fmt.Println("------------------------------------------- 传递数组")
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) //  传递数组的地址
	fmt.Printf("The sum of the array is: %f", x)

}

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

// 计算数组求和
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

/*

1、数组具有相同的唯一类型（整型，字符串，自定义等）的，并且长度固定，数组长度最大为2GB
2、数组长度必须是一个常量表达式，非负整数，而且长度也是数组的一部分，长度不同，数组类型也不同（ [5] int 和 [10] int 是属于不同类型的）
3、数组元素可以通过 索引（位置）来读取（或者修改），数组索引下标从0开始，数组第一个元素arr[0]，最后一个元素arr1[len(arr1)-1]
4、声明数组 var identifier [len]type 例如 var arr1 [5]int，声明数组时所有的元素都会被自动初始化为该类型的零值
5、数组是 可变的，可以对数组元素重新赋值，arr[i] = value，“可变”，指的是数组内容可以变化，长度不能变化
6、可以采用for循环，for-range循环 ,可以对数组进行初始化、依次处理对应元素
7、Go 语言中的数组是一种 值类型
8、var arr1 = new([5]int) 创建 => arr1 的类型是 *[5]int                var arr2 [5]int 创建 => arr2 的类型是 [5]int
9、数组常量 的方法来初始化数组，var arrAge = [5]int{18, 20, 15, 22, 16}  或者 var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
10、多维数组，内部数组总是长度相同的。Go 语言的多维数组是矩形式的
11、把一个大数组传递给函数会消耗很多内存，可以借助指针或者切片进行解决
*/
