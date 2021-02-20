/*
 * @Descripttion: json
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 13:56:02
 * @LastEditTime: 2021-02-20 14:34:10
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	page   int
	Fruits []string
}

type Response2 struct {
	// 可以指定json转化后 字段标签
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	fmt.Println("=============================【 json 编码 】======================================")
	// t := Marshal(value)
	// string(t)

	//=============================== 基本类型 json转化
	println("=======================[基本类型]")

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(20.21)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//=============================== 切片 json转化
	println("========================[切片]")
	slcD := []string{"apple", "peach", "pear"}

	fmt.Println("Original :", slcD)
	slcB, _ := json.Marshal(slcD)
	fmt.Println("JSON     :", string(slcB))

	//=============================== map json转化
	println("========================[map]")
	mapD := map[string]int{"apple": 5, "banana": 7}

	fmt.Println("Original :", mapD)
	mapB, _ := json.Marshal(mapD)
	fmt.Println("Json     :", string(mapB))

	//=============================== struc json转化
	println("========================[结构体]")
	// json编码仅导出可以编码的字段 ，也就是开头大写的
	// json编码会使用默认字段名作为json的key
	res1D := &Response1{
		page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	fmt.Println("Original :", res1D)
	res1B, _ := json.Marshal(res1D)
	fmt.Println("Json     :", string(res1B))

	// respons2 中跟1 不同的就是 指定默认编码字段 Page   int      `json:"page"`
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	fmt.Println("Original :", res2D)
	res2B, _ := json.Marshal(res2D)
	fmt.Println("Json     :", string(res2B))

	fmt.Println("=============================【 json 解码 】======================================")
	// Unmarshal([]byte, &type)

	// 需要解码的数据结构类型
	byt := []byte(`{"num":6.66  , "strs": ["A", "B" ]}`)
	// 我们需要找到一个可以存放解码数据的变量
	// 定义一个map key为字符串 value为任意类型
	var dat map[string]interface{}

	// 解码过程中 判断是否出错
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 当然为了使用解码中map的值，我们需要将其进行适当的类型转化
	// 将 num => float64
	fmt.Printf("%t \n", dat["num"])
	num := dat["num"].(float64)
	fmt.Printf("%t \n", num)

	// 访问嵌套的值 进行一系列的转化
	fmt.Printf("%t \n", dat["strs"])
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Printf("%t \n", str1)

	// 可以直接解码到我们自定义的的数据类型中
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	// 解码
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 直接输出
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
