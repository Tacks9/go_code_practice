/*
 * @Descripttion: 文件读写
 * @Date: 2020-12-08 17:01:54
 * @LastEditTime: 2020-12-09 16:01:34
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("file-reading-and-writing -------------------------文件读写")

	// 读取某个文件，按照行读取
	// readFile("buf")
	// readFile("line")

	// 复制一个文件内容到另一个文件中
	// cpFile()

	// 写入文件
	writeFile()
}

/*
1、文件使用指向 os.File 类型的指针来表示的，也叫做文件句柄
2、标准输入 os.Stdin 和标准输出 os.Stdout，他们的类型都是 *os.File
3、defer inputFile.Close() 程序退出前关闭该文件
4、bufio.NewReader 来获得一个读取器变量
5、读取器可以一行一行读取，比如ReadString('\n')、ReadBytes('\n')、ReadLine() 方法来实现相同的功能
6、Unix 和 Linux 的行结束符是 \n，而 Windows 的行结束符是 \r\n
7、在使用 ReadString 和 ReadBytes 方法的时候，我们不需要关心操作系统的类型，直接使用 \n 就可以了
*/
func readFile(typeRead string) {
	// 获取当前目录
	nowPath, _ := os.Getwd()
	println("当前目录:" + nowPath)

	// 【1】打开一个文件描述符 os
	inputFile, inputError := os.Open(nowPath + "/GoCode/chapter_7_read_write_data/02_test.txt")
	// 如果有错误
	if inputError != nil {
		fmt.Println(inputError)
		fmt.Printf("Does the file exist?\n")
		return
	}
	// 【2】确保在程序退出前关闭该文件
	defer inputFile.Close()

	// 【3】读取器 bufio
	inputReader := bufio.NewReader(inputFile)

	if typeRead == "buf" {
		// 开启缓存读取
		buf := make([]byte, 1024)
		for {
			// 【4】 读数据
			n, err := inputReader.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if 0 == n {
				break
			}
			fmt.Println(string(buf[0:n]))

		}
	} else {
		// 按照行读取
		for {
			// 【4】 读数据
			inputString, readerError := inputReader.ReadString('\n')
			fmt.Printf(":> %s ", inputString)
			// io.EOF = true
			if readerError == io.EOF {
				return
			}
		}

	}

}

/*
1、ioutil.ReadFile() 【 一次性读取文件 】
	func ReadFile(filename string) ([]byte, error)
2、ioutil.WriteFile(outFile, buf, 0644) 【 WriteFile 向文件 filename 中写入数据 data 】
	func WriteFile(filename string, data []byte, perm os.FileMode) error
	如果文件不存在，则以 perm 权限创建该文件
	如果文件存在，os.FileMode 不同使用不同的规则
*/
func cpFile() {
	// 获取当前目录
	nowPath, _ := os.Getwd()
	println("当前目录:" + nowPath)

	inputFile := nowPath + "/GoCode/chapter_7_read_write_data/02_cp.txt"
	// 可以不存在
	outFile := nowPath + "/GoCode/chapter_7_read_write_data/02_cp_copy.txt"

	// （ioutil.ReadFile() 方法，该方法第一个返回值的类型是 []byte）
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	// 写入另一个文件
	err = ioutil.WriteFile(outFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}

}

/*
[写文件]

*/
func writeFile() {
	// 【1】 以可写的方式打开一个文件  os.OpenFile
	// 【2】 确保在程序退出前关闭文件  Close
	// 【3】 写入器 bufio            bufio.NewWriter
	// 【4】 写入缓冲区              WriteString
	// 【5】 刷新缓冲区              Flush

	// 获取当前目录
	nowPath, _ := os.Getwd()
	println("当前目录:" + nowPath)

	// 可以不存在
	fileName := nowPath + "/GoCode/chapter_7_read_write_data/02_write_file.txt"

	// 【1】打开文件 (只写模式打开文件 只写模式打开文件)
	// 在读文件的时候，文件的权限是被忽略的，所以在使用 OpenFile 时传入的第三个参数可以用 0。
	// 而在写文件时，不管是 Unix 还是 Windows，都需要使用 0666。
	outFile, outErr := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if outErr != nil {
		fmt.Println(outErr)
		return
	}
	// 【2】关闭文件
	defer outFile.Close()
	// 【3】写入器
	outWriter := bufio.NewWriter(outFile)
	outString := "hello,world\n"
	// 【4】写
	for i := 0; i < 10; i++ {
		// 将字符串写入缓冲区，写 10 次
		outWriter.WriteString(outString)
	}
	// 【5】更新缓冲区
	// 缓冲区的内容紧接着被完全写入文件
	outWriter.Flush()

}
