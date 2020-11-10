package main

import (
	"fmt"
)

func main() {
	fmt.Println("switch structure")

	var num1 int = 100

	// case与待比较的变量值都是 int类型
	switch num1 {
	// 同时测试多个可能符合条件的值，使用逗号分割它们
	case 98, 99:
		// 每个case后的语句不需要{}括号括起来
		fmt.Println("It's equal to 98")
		fallthrough
		// 不需要特别使用 break 语句来表示结束
	case 100:
		fmt.Println("It's equal to 100")
		// fallthrough还希望继续执行后续分支的代码
		fallthrough
	default:
		fmt.Println("default ...")
	}

	// case与待比较的变量值都是 bool类型
	// 省略bool值
	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 <= 100:
		fmt.Println("Number is between 0 and 100")
	default:
		fmt.Println("Number is 10 or greater")
	}

	// 条件初始化语句 平行赋值
	switch a, b := 1, 2; {
	case a > b:
		fmt.Println("a>b")
	case a < b:
		fmt.Println("a<b")
	case a == b:
		fmt.Println("a==b")
	default:
		fmt.Println("a ？ b")
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}

/*
	1、switch中条件的变量类型不被局限于常量或整数，但必须是相同的类型
	2、一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块
	3、case语句后不需要使用花括号将多行语句括起来
	4、如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字
	5、不需要特别使用 break 语句来表示结束
	6、可以省略swtich后面变量，而进行比较布尔值
	7、switch 可以后面跟上平行赋值语句
*/
