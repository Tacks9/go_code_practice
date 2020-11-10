package main

import (
	"fmt"
)

func main() {
	fmt.Println("break  and  continue")

	// 标签LABEL1
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				// continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置
				continue LABEL1
			}
			// 当 j==4 和 j==5 的时候，没有任何输出
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	a := 1
	b := 9
	//  应当只使用正序的标签（标签位于 goto 语句之后），但注意标签和 goto 语句之间不能出现定义新变量的语句
	goto TARGET
	a++ // 会被跳过
TARGET:
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)

}

/*
1、for、switch 或 select 语句都可以配合标签（label）形式的标识符使用
2、标签的名称是大小写敏感的，为了提升可读性，一般建议使用全部大写字母
3、应当只使用正序的标签（标签位于 goto 语句之后）
*/
