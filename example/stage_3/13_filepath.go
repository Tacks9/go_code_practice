/*
 * @Descripttion: 文件路径
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 16:22:32
 * @LastEditTime: 2021-02-20 17:43:04
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var P = fmt.Println

// 定义一下错误函数
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("=============================【 文件路径 】======================================")

	// Join可以跨操作系统 接收任意数量参数 构建路径
	path1 := filepath.Join("dir1", "dir2", "filename")
	P("path1: ", path1)

	// 尽量使用Joins来进行拼接，而避免使用/，虽然/也会进行自动去除
	// P(filepath.Join("dir1//", "demo"))
	// P(filepath.Join("dir1//", "../../", "demo"))

	// 上一层目录
	P("Dir(path1): ", filepath.Dir(path1))
	// 文件
	P("Base(path1): ", filepath.Base(path1))

	// 判断是否是绝对路径
	P(filepath.IsAbs("dir/file"))
	P(filepath.IsAbs("/dir/file"))
	P(filepath.IsAbs(filepath.Join("/D", "work")))

	// 扩展名
	fileName := "config.json"
	// 获取文件扩展名
	ext := filepath.Ext(fileName)
	// 纯文件名 不要后缀
	name := strings.TrimSuffix(fileName, ext)
	P("[file]", fileName, "[Ext]", ext, "[Name]", name)

	// 相对路径
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	fmt.Println("=============================【 目录 】======================================")
	//  创建一个目录
	dirErr := os.Mkdir("subdir", 0755)
	check(dirErr)

	// 延迟操作defer 在函数最后记得删除 目录
	defer os.RemoveAll("subdir")

	// 创建临时文件的函数
	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile(filepath.Join("subdir", "file1.txt"))

	// `ReadDir` 列出目录的内容，返回一个 `os.FileInfo` 类型的切片对象。
	clist, dirErr := ioutil.ReadDir("subdir")
	check(dirErr)

	// 遍历一个目录
	for _, item := range clist {
		// 文件 是否是目录
		P(" ", item.Name(), item.IsDir())
	}

	// 切换工作目录
	dirErr = os.Chdir("subdir")
	check(dirErr)

	// `cd` 回到最开始的地方。  为什么要回去，因为你最后还是要删除 subdir目录
	err = os.Chdir("..")
	check(err)

	// 读取一下当前目录
	clist, dirErr = ioutil.ReadDir(".")
	check(dirErr)

	for _, entry := range clist {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 递归访问目录
	filepath.Walk("./GoCode/example", visit)

	fmt.Println("=============================【 临时目录/文件  ioutil.TempFile 】======================================")
	// `ioutil.TempFile` 会在操作系统的默认位置下创建该文件
	ff, err := ioutil.TempFile("", "sample")
	check(err)

	// 打印临时文件名
	fmt.Println("Temp file name:", ff.Name())

	// defer 删除该文件。
	// 尽管操作系统会自动在某个时间清理临时文件，但手动清理是一个好习惯。
	defer os.Remove(ff.Name())

	// 我们可以向文件写入一些数据。
	_, err = ff.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := ioutil.TempDir("", "sampledir")
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// 现在，我们可以通过拼接临时目录和临时文件合成完整的临时文件路径，并写入数据。
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

// walk遍历到每一个目录或者文件，都会调用访问
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	P(" ", p, info.IsDir())
	return nil
}
