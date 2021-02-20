/*
 * @Descripttion: URL 解析 net/url
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 15:52:31
 * @LastEditTime: 2021-02-20 16:09:09
 */
package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("=============================【 Url 解析 】======================================")
	// URL 统一资源定位符 认证信息、主机、端口、路径、查询参数、片段
	// s := "postgres://user:pass@host.com:5432/path?k=v#f"
	s := "https://learnku.com:80/docs/gobyexample/2020/url-analysis/6306"
	fmt.Println(s)

	// 解析一个url 并确保没有出错
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
	// 直接访问 Scheme
	fmt.Println(u.Scheme)

	// User包含 用户密码
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	pass, _ := u.User.Password()
	fmt.Println(pass)

	// Host包含 主机端口
	fmt.Println(u.Host)
	// 分割
	host := strings.Split(u.Host, ":")
	fmt.Println(host[0])
	fmt.Println(host[1])

	// 路径
	fmt.Println(u.Path)
	// 片段
	fmt.Println(u.Fragment)

	// 参数
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
}
