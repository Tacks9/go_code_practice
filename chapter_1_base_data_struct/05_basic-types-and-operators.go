package main

import (
	"fmt"
	"math/rand" // 数学函数包
	"time"      // 时间包
	"unicode"   // 字符unicode包
)

func main() {
	fmt.Println("Hello! Go World! this is basic types and operators! ")

	fmt.Println("------------------------------------------布尔型") // true

	var flag bool = true
	fmt.Println(flag)  // true
	fmt.Println(!flag) // false
	var a int = 10
	var b int = 11
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	// 短路效应 && ||
	// && 左边表达式已经为false就会直接结束 不会看后面的表达式
	// || 左边表达式已经为true就会直接结束 不会看后面的表达式
	fmt.Println(a == b && flag) // false
	fmt.Println(a == b || flag) // true

	fmt.Println("------------------------------------------整型") // true
	// 不同类型混合不能混合使用
	var one int
	var two int32
	one = 15
	// two = one + one// 编译错误
	two = int32(one + one) // 编译成功
	fmt.Printf("int is: %d\n", two)
	two = two + 5 // 5是常量可以直接编译
	fmt.Printf("int is: %d\n", two)

	fmt.Println("------------------------------------------字符")
	var ch byte = 65
	fmt.Printf("ch is: %c ! \n", ch)

	fmt.Println(unicode.IsLetter('A'))
	fmt.Println(unicode.IsSpace(' '))
	fmt.Println(unicode.IsDigit(1))

	fmt.Println("------------------------------------------位运算")

	// 位运算
	// &
	m := 1 & 1 // 都1 才是1
	n := 1 & 0
	i := 0 & 1
	j := 0 & 0
	fmt.Printf("& 按位与 m=%b, n=%b, i=%b, j=%b \n", m, n, i, j)
	// |
	m = 1 | 1 // 有1就是1
	n = 1 | 0 // 有1就是1
	i = 0 | 1 // 有1就是1
	j = 0 | 0
	fmt.Printf("| 按位或 m=%b, n=%b, i=%b, j=%b \n", m, n, i, j)
	// ^
	m = 1 ^ 1 // 相同为0
	n = 1 ^ 0 // 不同为1
	i = 0 ^ 1
	j = 0 ^ 0
	fmt.Printf("^ 按位异或  m=%b, n=%b, i=%b, j=%b \n", m, n, i, j)
	// 位左移 <<  bitP << n
	// 向左移动 n 位，右侧空白部分使用 0 填充
	// 如果 n 等于 2，则结果是 2 的相应倍数，即 2 的 n 次方
	m = 1 << 1
	n = 1 << 2
	i = 1 << 4
	j = 1 << 5
	fmt.Printf("<< 位左移  m=%d, n=%d, i=%d, j=%d \n", m, n, i, j)

	// 位右移
	m = 1 >> 1
	n = 1 >> 2
	i = 1 >> 4
	j = 1 >> 5
	fmt.Printf(">> 位右移  m=%d, n=%d, i=%d, j=%d \n", m, n, i, j)

	type ByteSize float64
	const (
		_           = iota // 通过赋值给空白标识符来忽略值
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)
	fmt.Printf("  KB=%g, MB=%g, GB=%g, TB=%g \n", KB, MB, GB, TB)

	fmt.Println("------------------------------------------算数运算符")
	numa := 10
	numb := 2
	var numc int = 0
	numc = numa + numb
	fmt.Printf("  numa + numb = %d\n", numc)
	numc = numa - numb
	fmt.Printf("  numa + numb = %d\n", numc)
	numc = numa * numb
	fmt.Printf("  numa + numb = %d\n", numc)
	numc = numa / numb
	fmt.Printf("  numa + numb = %d\n", numc)
	numc = numa % numb
	fmt.Printf("  numa + numb = %d\n", numc)
	// numc = numa++ // 这个会报错
	// fmt.Printf("  numa ++ = %d\n", numc)
	numa++
	numb++
	numc = numa + numb
	fmt.Printf("  numa + numb = %d\n", numc)

	fmt.Println("------------------------------------------类型别名")
	type TZ int
	var pa, pb TZ = 3, 4
	pc := pa + pb
	fmt.Printf("pc has the value: %d", pc) // 输出：c has the value: 7

	fmt.Println("------------------------------------------随机数")

	for i := 0; i < 6; i++ {
		intNum := rand.Int()
		fmt.Printf("%d / ", intNum)
	}
	fmt.Println()
	for i := 0; i < 6; i++ {
		intNum := rand.Intn(8) // rand.Intn 返回介于 [0, n) 之间的伪随机数
		fmt.Printf("%d / ", intNum)
	}
	fmt.Println()

	timens := int64(time.Now().Nanosecond()) // 当前时间的纳秒级数字
	rand.Seed(timens)                        // 随机数的生成种子

	for i := 0; i < 6; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32()) // rand.Float32 和 rand.Float64 返回介于 [0.0, 1.0) 之间的伪随机数
	}

}

/*

布尔类型 Bool
	布尔型的值只可以是常量 true 或者 false
	%t 来表示你要输出的值为布尔型
数字类型 int float
	1、int、uint、uintptr、（有符号整数）int8、int16、int32、int64、（无符号整数）uint8、uint16、unt32、uint64
	2、浮点型只有，float32（精确到小数点后 7 位）、float64（float64 精确到小数点后 15 位），记住！！！没有float也没有double
	3、整型的零值为0
	4、浮点型的零值为0.0
	5、不允许不同类型之间的混合使用，但是允许常量之间的混合使用
	6、%d 用于格式化整数 、%X 用于格式化 16 进制表示的数字、%g 用于格式化浮点型、%f 输出浮点数、%b 是用于表示位的格式化标识符
字符类型
	1、字符只是整数的特殊用例。byte 类型是 uint8 的别名
	2、只占用 1 个字节的传统 ASCII 编码的字符
	3、var ch byte = 'A' 字符使用单引号括起来
	4、%c 用于表示字符, %v 或 %d 会输出用于表示该字符的整数
	5、包 unicode 包含了一些针对测试字符的非常有用的函数 （判断是否是字母 unicode.IsLetter(ch)，是否包含空白字符unicode.IsSpace(ch)）
逻辑运算符
	1、可以通过逻辑运算符来产生布尔值，如!、&&、||
	2、两个类型相同的值可以使用相等 == 或者不等 !=运算符来进行比较，如果类型不同需要显示强制类型转化
	3、比较运算符 <、<=、>、>=
	4、逻辑运算符的结果都是布尔值
算数运算符
	1、+、-、*、/、%、+=、-=、*=、/=、%=、++、--
	2、/ 对于整数运算而言，结果依旧为整数，例如：9 / 4 -> 2
	3、字符串拼接时允许使用对运算符 + 的重载
	4、取余运算符只能作用于整数：9 % 4 -> 1
	5、带有 ++ 和 -- 的只能作为语句，而非表达式，n = i++ 这种写法是无效的
运算符的优先级
	最好用小括号来更清楚显示优先级
类型别名
	1、`type TZ int` 中，TZ 就是 int 类型的新名称
	2、类型别名得到的新类型并非和原类型完全相同，新类型不会拥有原类型所附带的方法
其他
	"math/rand" // 数学函数包
	"time"		// 时间包
	"unicode"	// 字符unicode包

*/
