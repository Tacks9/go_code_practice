package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello! Go World! this is strings strconv packages ! ")

	fmt.Println("---------------------------------------------------string 包")

	var str string = "This is Go Lang end"
	fmt.Println(strings.HasPrefix(str, "Th"))             // true
	fmt.Println(strings.HasSuffix(str, "End"))            // false
	fmt.Println(strings.Contains(str, "end"))             // true
	fmt.Println(strings.Index(str, "is"))                 // 2
	fmt.Println(strings.LastIndex(str, "is"))             // 5
	fmt.Println(strings.Replace(str, "Go", "GoLand", 10)) // This is GoLand Lang end
	fmt.Println(strings.Count(str, "is"))                 // 2
	fmt.Println(strings.Repeat(str, 2))                   // This is Go Lang endThis is Go Lang end
	fmt.Println(strings.ToLower(str))                     // this is go lang end
	fmt.Println(strings.ToUpper(str))                     // THIS IS GO LANG END

	str = " My name is Tacks "
	fmt.Println(strings.TrimSpace(str))          // My name is Tacks
	fmt.Println(strings.Trim(str, "My name is")) // Tacks
	fmt.Println(strings.Fields(str))             // [My name is Tacks]
	fmt.Println(strings.Split(str, " "))         // [ My name is Tacks ]

	fmt.Println("---------------------------------------------------string 包 Fields string按照空格拆成slice")

	str1 := "The quick brown fox jumps over the lazy dog"
	str2 := strings.Fields(str1)
	fmt.Printf("str1: %s\n", str1)
	// Fields 按照空格分隔
	fmt.Printf("Splitted in slice: %v\n", str2)
	for _, val := range str2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	fmt.Println("---------------------------------------------------string 包 Split按照特定分隔符 划分字符串")

	fmt.Println()
	str3 := "GO1|The ABC of Go|25"
	// Split 按照特定分割符分割
	str4 := strings.Split(str3, "|")
	fmt.Printf("Splitted in slice: %v\n", str4)
	for _, val := range str4 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println("---------------------------------------------------string 包  Join合并")

	fmt.Println()
	// Join 将slice按照; 进行合并成 slice
	str5 := strings.Join(str4, ";")
	fmt.Printf("str4 joined by ;: %s\n", str5)

	fmt.Println("---------------------------------------------------strconv 包  类型转化")

	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig) // string=>int
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an) // int=>string
	fmt.Printf("The new string is: %s\n", newS)

}

/*

strings包
	1、HasPrefix 判断字符串 s 是否以 prefix 开头 `strings.HasPrefix(str, "Th")`
	2、HasSuffix 判断字符串 s 是否以 suffix 结尾 `strings.HasSuffix(str, "End")`
	3、Contains 判断字符串 s 是否包含 substr  `strings.Contains(s, substr string) bool`
	4、Index 获取子字符串或字符在父字符串中出现的位置 `strings.Index(s, str string) int`
	5、LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引 `strings.LastIndex(s, str string) int`
	6、Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new `strings.Replace(str, old, new, n) string`
	7、Count 用于计算字符串 str 在字符串 s 中出现的非重叠次数 `strings.Count(s, str string) int`
	8、Repeat 用于重复 count 次字符串 s 并返回一个新的字符串 `strings.Repeat(s, count int) string`
	9、ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符 `strings.ToLower(s) string`
	10、ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符 `strings.ToUpper(s) string`
	11、strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号
	12、strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉
	13、只想剔除开头或者结尾的字符串，则可以使用 TrimLeft 或者 TrimRight
	14、strings.Fields(s) 利用空白作为分隔符将字符串分割为若干块，并返回一个 slice
	15、strings.Split(s, sep) 自定义分割符号对字符串分割，返回 slice
	16、Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串 `strings.Join(sl []string, sep string) string`
strconv
数字类型转换到字符串
	1、返回数字 i 所表示的字符串类型的十进制数 `strconv.Itoa(i int) string`
	2、strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string
字符串类型转换为数字类型（并不总是可能的）
	1、strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型
	2、strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型
*/
