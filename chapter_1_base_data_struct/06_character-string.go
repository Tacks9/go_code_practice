package main

import (
	"fmt"
	"unicode/utf8" // utf8的编码格式
)

func main() {
	fmt.Println("Hello! Go World! this is character string! ")

	fmt.Println("`This is a raw string ` 中的 ` ")
	fmt.Println("This is a \t string \n 中的 \n ")

	var s1 string = "My name is "
	var s2 string = "tacks"
	fmt.Printf("%s\n", s1+s2) // 字符串拼接符

	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))

	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2))

}

/*
字符串
	1、Go 中的字符串也可能根据需要占用 1 至 4 个字节
	2、字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容
	3、字符串是字节的定长数组。
	4、空字符串 ""
	5、获取字符串长度 len(str)
	6、字符串内容可以通过标准索引进行获取 如字符串 str 的第 1 个字节：str[0] ，最后 1 个字节：str[len(str)-1]
	7、字符串中某个字节的地址的行为是非法的，例如&str[i]
	8、字符串连接符 +
	9、拼接字符更好的办法是使用函数 strings.Join()

转义字符
	1、该类字符串使用双引号括起来，其中的相关的转义字符将被替换
	2、\n换行符  \r回车符  \t tab 键  \u Unicode 字符  \\ 反斜杠自身
非转义字符
	1、使用反引号括起来
	2、`This is a raw string \n` 中的 `\n\` 会被原样输出。
*/
