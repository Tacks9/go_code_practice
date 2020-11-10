package main

import (
	"fmt"
	"time" // 日期时间包
)

func main() {
	fmt.Println("Hello! Go World! this is time and date packages ! ")

	fmt.Println("---------------------------------------------------time 包")

	timeNow := time.Now()
	fmt.Println(time.Now())                 // 2020-11-09 16:20:15.4638986 +0800 CST m=+0.004880401
	fmt.Println(timeNow.UTC())              // UTC
	fmt.Println(timeNow.Year())             // 年份
	fmt.Println(timeNow.Day())              // 小时
	fmt.Println(timeNow.Minute())           // 分钟
	fmt.Println(timeNow.Format(time.ANSIC)) // Mon Nov  9 16:27:28 2020
}

/*

time
	1、提供显示和测量时间和日期的功能函数的包
	2、当前时间 time.Now()
	3、当前天   t.Day()
	4、当前分钟 t.Minute()
*/
