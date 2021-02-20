/*
 * @Descripttion:  base64 编码 (encoding/base64)
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 16:09:55
 * @LastEditTime: 2021-02-20 16:17:41
 */
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("=============================【 base64 编码处理 】======================================")
	// 标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在 稍许不同（后缀为 + 和 -），但是两者都可以正确解码为 原始字符串

	fmt.Println("【标准Base64】")
	// 待编码的字符串
	data := "Galaxy-Park-9 !?$*&()'-=@~"

	// 编码需要 []byte 类型
	sEn := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEn)

	// 解码可能出错，所以要进行判断
	sDe, _ := base64.StdEncoding.DecodeString(sEn)
	fmt.Println(string(sDe))

	fmt.Println("【URL兼容】")
	uEn := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEn)
	uDe, _ := base64.URLEncoding.DecodeString(uEn)
	fmt.Println(string(uDe))

}
