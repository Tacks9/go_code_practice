/*
 * @Descripttion: for range 遍历 不同的数据类型
 * @Author: tacks321@qq.com
 * @Date: 2021-02-09 15:04:14
 * @LastEditTime: 2021-02-09 15:43:19
 */

package main

import "fmt"

func main() {
	fmt.Println("=================for range 遍历 =======================")

	// ==============================  数组
	varArr := [5]int{1, 3, 5, 7, 8}
	fmt.Println("[varArr] :", varArr)
	// 数组求和
	sum := 0
	// for 索引, 值 := range 数组 {}
	for _, val := range varArr {
		sum = val + sum
	}
	fmt.Println("array get sum: ", varArr, sum)

	// ==============================  切片
	varSlice := varArr[0:3]
	fmt.Println("[varSlice] :", varSlice, "len:", len(varSlice), "cap:", cap(varSlice))
	// for 索引, 值 := range 切片 {}
	for key, val := range varSlice {
		fmt.Print(key, " ")
		if key == 3 {
			fmt.Println("key:", key, "val:", val)
		}
	}
	fmt.Println("")

	// ============================== map
	varMap := map[string]string{"Name": "Tacks", "Age": "18"}
	for key, val := range varMap {
		fmt.Printf("%s===>%s\n", key, val)
	}

	// ============================== string
	varStr := "Tacks"
	for i, ascii := range varStr {
		fmt.Println(i, ascii)
	}
}
