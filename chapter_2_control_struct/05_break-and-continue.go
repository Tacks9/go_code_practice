package main

import (
	"fmt"
)

func main() {
	fmt.Println("break  and  continue")

	// --------------- 采用break 来作为for循环终止条件
	var i int = 5
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i <= 0 {
			break
		}
	}
	// -------------------- break 只会退出最内层的循环
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			print(j)
		}
		print("  ")
	}
	println("")
	// ----------------------- continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 关键字 continue 只能被用于 for 循环中
		}
		print(i)
		print(" ")
	}

}

/*
1、break 语句的作用结果是跳过整个代码块，执行后续的代码（ 只会退出最内层的循环 ）
2、可以将for循环改成类似while循环，加上break的判断条件，从而控制循环的终止
3、continue 忽略剩余的循环体而直接进入下一次循环的过程 （ 关键字 continue 只能被用于 for 循环中 ）
*/
