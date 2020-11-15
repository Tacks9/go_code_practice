package main

import "fmt"

func main() {
	fmt.Println("test-key-values-for-the-existence-and-deletion-of-elements -------------------------测试键值对是否存在及删除元素")

	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["NewDelhi"] = 55
	map1["BeiJing"] = 20
	map1["Washington"] = 25
	// 判断key是否存在
	value, isPresent = map1["BeiJing"]
	if isPresent {
		fmt.Printf("map1 key=BeiJing value=%d\n", value)
	} else {
		fmt.Printf("map1 key=BeiJing value=not")
	}
	fmt.Println("---")

	value, isPresent = map1["Paris"]
	fmt.Printf("map1 key=Paris value=%d isPresent=%t", value, isPresent)

	fmt.Println("---")

	// 删除元素
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("map1 key=Washington value=%d\n", value)
	} else {
		fmt.Printf("map1 key=Washington value=not")
	}

}

/*


1、如果 map 中不存在 key1，val1 就是一个值类型的空值
2、测试键值对是否存在 val1, isPresent = map1[key1]
3、isPresent 返回一个 bool 值：如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，并且 isPresent 为 true
4、只是想判断某个 key 是否存在而不关心它对应的值到底是多少  _, ok := map1[key1]
5、从 map1 中删除 key1 。可以使用delete(map1, key1)，如果 key1 不存在，该操作不会产生错误

*/
