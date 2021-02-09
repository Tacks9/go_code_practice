/*
 * @Descripttion: map 关联数组 引用类型 采用make声明 长度可伸缩
 * @Author: tacks321@qq.com
 * @Date: 2021-02-09 11:47:38
 * @LastEditTime: 2021-02-09 15:02:34
 */

package main

import "fmt"

func main() {
	fmt.Println("===================maps=========================")

	// make创建一个map
	map1 := make(map[string]int)

	// set
	map1["a"] = 1
	map1["b"] = 2

	// 打印
	fmt.Println("[map1] ALL:", map1)
	// get
	fmt.Println("[map1] a:", map1["a"])
	// len
	fmt.Println("[map1] Len:", len(map1))
	// del
	delete(map1, "a")
	fmt.Println("[map1] ALL:", map1)

	// _空白标识符用法
	_, is_ok := map1["a"]
	fmt.Println("[map1] a is:", is_ok)

	// 声明一个map并初始化
	map2 := map[string]int{"one": 1, "two": 2}
	fmt.Println("[map2] All:", map2)

	//
	var map3 map[string]int
	var map4 map[string]int

	// 这个地方相当于make？
	map3 = map[string]int{"key1": 1, "key2": 2}
	// 引用赋值
	map4 = map3

	// 更改 map4，那么map3同时也会改变
	map4["key4"] = 4

	fmt.Println("[map3] All:", map3)
	fmt.Println("[map4] All:", map4)
	// 打印不存在的key，会被默认给成零值
	fmt.Println("[map4] a:", map4["a"])

}
