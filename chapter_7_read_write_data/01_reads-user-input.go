/*
 * @Descripttion: 读取用户输入相关
 * @Author: tacks321@qq.com
 * @Date: 2020-12-01 18:03:16
 * @LastEditTime: 2020-12-08 17:01:58
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

// var (
// 	firstName string
// 	lastName  string
// 	s         string
// 	i         int
// 	f         float32
// 	input     = "56.12 / 5212 / Go"
// 	format    = "%f / %d / %s"
// )
// inputReader 是一个指向 bufio.Reader 的指针
var inputReader *bufio.Reader
var input string
var err error

func main() {
	fmt.Println("read user input -------------------------读取用户的输入")

	// fmt.Println("Please enter your full name: ")
	// // 从控制台读取
	// fmt.Scanln(&firstName, &lastName)
	// fmt.Printf("Hi %s %s!\n", firstName, lastName)

	// fmt.Sscanf(input, format, &f, &i, &s)
	// fmt.Println("From the string we read:", f, i, s)

	// inputReader 是一个指向 bufio.Reader 的指针.
	// 函数返回一个新的带缓冲的 io.Reader 对象 。它将从指定读取器（例如 os.Stdin）读取内容
	inputReader = bufio.NewReader(os.Stdin)

	fmt.Println("Please enter some input: ")
	// 从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s \n", input)
	}

}

/*


 */
