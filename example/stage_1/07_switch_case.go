/*
 * @Descripttion: switch/case 多分支快捷操作
 * @Author: tacks321@qq.com
 * @Date: 2021-02-07 17:07:46
 * @LastEditTime: 2021-02-07 17:48:20
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=========================switch case 分支===========================")
	i := 2
	// 简单用法
	switch i {
	case 1:
		fmt.Println("太厉害了 No.1")
	case 2:
		fmt.Println("第二名也不错")
	case 3:
		fmt.Println("下次一定要拿第一")
	}

	// 使用逗号，使用default默认选项
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("This is Weekend")
	default:
		fmt.Println("It is a weekday")
	}

	// 不带表达式的switch 转化成if else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is Before noon")
	default:
		fmt.Println("It is after noon")
	}

	// 封装一个函数
	whatAmI := func(i interface{}) {
		switch iType := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Printf("%T", iType)
		}
	}

	whatAmI(9)
	whatAmI("Tacks")
	whatAmI(true)

	// 关键词 fallthrough
	// 如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字
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
