/*
 * @Descripttion: 随机数 rand
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 15:39:21
 * @LastEditTime: 2021-02-20 15:50:59
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=============================【 随机数生成 rand 】======================================")
	// 默认情况下种子是一定的 ，所以会返回一样的值
	// 随机返回一个整数 0 ~ 100
	fmt.Println(rand.Intn(100))
	// 随机返回一个浮点数 0.0 ~ 1.0
	fmt.Println(rand.Float64())
	// 可以生成其他范围的随机浮点数 5.0 ~ 10.0
	fmt.Println((rand.Float64() * 5) + 5)

	fmt.Println("【变化种子】")
	// 变化种子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// 每次都会生成不同的值
	fmt.Println(r1.Intn(100))

	fmt.Println("【给定种子】")
	// 给定种子
	s2 := rand.NewSource(99)
	r2 := rand.New(s2)

	fmt.Println(r2.Intn(99))
}
