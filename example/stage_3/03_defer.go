/*
 * @Descripttion: 延迟调用  defer 可以做一些函数结束前的清理工作
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 09:45:36
 * @LastEditTime: 2021-02-20 09:57:41
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=============================【 延迟调用 defer】======================================")

	fd := createFile("defer.txt")
	// 在main最后执行完毕的时候，会进行关闭 相当于最后可以做一些清理工作
	// 在创建文件后，紧跟着延迟关闭文件，防止最后没有关闭
	defer cloaseFile(fd)

	writeFile(fd)

}

// 创建文件
func createFile(p string) *os.File {

	fmt.Println("Creating...")
	fd, err := os.Create(p)
	// 创建失败错误处理
	if err != nil {
		panic(err)
	}
	return fd
}

// 写入内容
func writeFile(fd *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(fd, "data")
}

// 关闭文件
func cloaseFile(fd *os.File) {
	fmt.Println("Closed")
	fd.Close()
}
