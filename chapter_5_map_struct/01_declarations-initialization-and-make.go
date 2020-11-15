/*
	2020/11/15 周日
*/
package main

import "fmt"

func main() {
	fmt.Println("This is declarations-initialization-and-make -------------------------声明、初始化、make")

	var mapList map[string]int
	var mapCopy map[string]int

	// 初始化map，类似数组一样初始化
	mapList = map[string]int{"one": 1, "two": 2}
	// make申请一个map。 引用类型 内存用 make 方法来分配 var map1 = make(map[keytype]valuetype)
	mapMake := make(map[string]float32)
	// 引用赋值  对 mapCopy 的修改也会影响到 mapList 的值
	mapCopy = mapList

	// map追加
	mapMake["key1"] = 4.5
	mapMake["key2"] = 3.1415926
	// 同时会改变 mapList
	mapCopy["two"] = 22

	fmt.Printf("Map mapList at \"one\" is: %d\n", mapList["one"])   // 1
	fmt.Printf("Map mapMake at \"key2\" is: %f\n", mapMake["key2"]) // 301415926
	fmt.Printf("Map mapCopy at \"two\" is: %d\n", mapCopy["two"])   // 22
	// 如果 map 中不存在 key1，val1 就是一个值类型的空值
	fmt.Printf("Map mapList at \"ten\" is: %d\n", mapList["ten"]) // 0 不存在的直接返回0

	fmt.Println("-------")

	// 不要使用 new，永远用 make 来构造 map
	mapNew := new(map[string]float32)
	// 如果你错误的使用 new () 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址
	// mapNew["key1"] = 4.5 // nvalid operation: mapNew["key1"] (type *map[string]float32 does not support indexing) 报错
	fmt.Println(mapNew) // &map[]

	fmt.Println("-------map的值是可以是任何类型的")
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
	fmt.Println("-------用切片作为 map 的值")
	// mp1 := make(map[int][]int) // value 定义为 []int 类型或者其他类型的切片

}

/*
1、声明一个map类型，var map1 map[keytype]valuetype ,例如var map1 map[string]int
2、没有被初始化的map变量的值为nil
3、声明的时候不要知道map的值，map是可以动态增长的，所以不需要跟数组那样设置容量
4、map可以根据新插入的key value 弹性伸缩，它不存在固定长度或者最大限制，但是也可以设置map的初始容量 capacity，例如 make(map[keytype]valuetype, cap)
5、key的类型，只能是string\int\float，数组切片这些不能作为key, 另外只包含内建类型的struct可以作为key,如果指针或者接口类型，可以通过Key()\Hash()方法获取唯一字符串key
6、value的类型，可以是任意的
7、map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节
8、根据key在map中找到value是非常快的，但是仍然比从数组和切片的索引中直接读取要慢 100 倍
9、map 也可以用函数作为自己的值
10、常用的 len(map1) 方法可以获得 map 中的 pair 数目，这个数目是可以伸缩的，因为 map-pairs 在运行时可以动态添加和删除。
11、如果 map 中不存在 key1，val1 就是一个值类型的空值


*/
