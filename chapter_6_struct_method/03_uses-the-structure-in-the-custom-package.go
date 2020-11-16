package main

import (
	"fmt"
	// 使用自定义包中的结构体
	"./structPack"
)

func main() {
	fmt.Println("This is uses-the-structure-in-the-custom-package -------------------------使用自定义包中的结构体")

	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.

	fmt.Printf("Mi1 = %d\n", struct1.Mi1) // Mi1 = 10
	fmt.Printf("Mf1 = %f\n", struct1.Mf1) // Mf1 = 16.000000
}
