package main

import (
	"fmt"
)

func main() {
	fmt.Println("切片---------------------------------slice")

	var arr1 [6]int              // 数组 相关数组
	var slice1 []int = arr1[2:5] // 切片 // item at index 5 not included!
	// 初始化数组
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	fmt.Println("----------------------------------- 数组初始化，切片初始化")
	fmt.Println(arr1)
	fmt.Println(slice1)
	// 打印切片
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))       // 数组长度
	fmt.Printf("The length of slice1 is %d\n", len(slice1))   // 切片长度
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1)) // 切片容量

	// 切片扩展大小
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of slice1 is %d\n", len(slice1))   // 切片长度
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1)) // 切片容量

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// 打印出来是整型
	fmt.Println(b[1:4]) // o l a
	fmt.Println(b[:2])  // g o
	fmt.Println(b[2:])  // l a n g
	fmt.Println(b[:])   // g o l a n g

	fmt.Println("----------------------------------- 切片传递给函数")
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(arr[:])) // 参数为切片

	fmt.Println("----------------------------------- make 创建切片")
	sliceMake1 := make([]int, 10) // make() 函数来创建一个切片 同时创建好相关数组 切片跟数组长度为10 类型为int

	// 这样分配一个有 10 个 int 值的数组，并且创建了一个长度为 5，容量为 10 的 切片 ，该 切片 指向数组的前 5 个元素。
	sliceMake2 := make([]int, 5, 10) //  make() 函数创建大小10的int数组，切片大小为5

	sliceMake3 := new([10]int)[0:5] // new() 创建10大小int的数组，然后切片为0~5的
	for i := 0; i < len(sliceMake1); i++ {
		sliceMake1[i] = i
	}
	for i := 0; i < len(sliceMake2); i++ {
		sliceMake2[i] = i
	}
	for i := 0; i < len(sliceMake3); i++ {
		sliceMake3[i] = i
	}
	fmt.Println(sliceMake1, len(sliceMake2), cap(sliceMake2)) // [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(sliceMake2, len(sliceMake2), cap(sliceMake2)) // [0 1 2 3 4]
	fmt.Println(sliceMake3, len(sliceMake3), cap(sliceMake3)) // [0 1 2 3 4] 5 10

	fmt.Println("----------------------------------- new () 和 make () 的区别")
	// new (T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的内存地址：
	// 这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体
	// var p *[]int = new([]int) // *p == nil; with len and cap 0

	// make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel
	// var v []int = make([]int, 10, 50)

	fmt.Println("-----------------------------------  bytes包中的  Buffer类型")
	/*
		// Buffer 可以这样定义：var buffer bytes.Buffer
		var buffer bytes.Buffer
		// 通过 buffer.WriteString(s) 方法将字符串 s 追加到后面
		for {
			if s, ok := getNextString(); ok {
				buffer.WriteString(s) // 追加字符串
			} else {
				break
			}
		}
		// buffer.String() 方法转换为 string
		fmt.Println(buffer.String())
	*/
}

// 求和
func sum(arr []int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

/*

1、切片（slice）是对数组（相关数组）一个连续片段的引用，切片是一个引用类型
2、切片是一个 长度可变的数组。和数组不同的是，切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度
3、切片 s 来说该不等式永远成立：0 <= len(s) <= cap(s)
4、数组实际上是切片的构建块，多个切片表示数组同一片段，可以进行共享数据
5、一个切片在未初始化之前默认为 nil，长度为 0
6、切片初始化 slice1 是由数组 arr1 中某一片段进行切分出来的 ar slice1 []type = arr1[start:end]
7、切片初始化 一个长度为 5 的数组并且创建了一个相关切片 var x = []int{2, 3, 5, 7, 11}
8、切片在内存中的组织方式实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度以及切片容量
9、绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针
10、可以使用 make () 函数来创建一个切片 同时创建好相关数组 var slice1 []type = make([]type, len, cap) ，其中 cap 是可选参数
11、切片make创建 make([]int, 50, 100) 或者 make([]int, 50)
12、bytes 包和字符串包十分类似 它还包含一个十分有用的类型 Buffer
*/
