package main

import (
	"fmt"
	"reflect"
)

// 带标签的结构体
type TagType struct {
	field1 bool   "An Important Answer"
	field2 string "The Name Of The Thing"
	field3 int    "How Much There Are"
}

// 打印结构体标签
func refTag(tt TagType, ix int) {
	// 在一个变量上调用 reflect.TypeOf() 可以获取变量的正确类型
	ttType := reflect.TypeOf(tt)
	// 变量是一个结构体类型，就可以通过 Field 来索引结构体的字段
	ixField := ttType.Field(ix)
	// 然后就可以使用 Tag 属性。
	fmt.Printf("%v %v\n", ixField, ixField.Tag)
}
func main() {
	fmt.Println("This is tag-structure -------------------------带标签的结构体")

	tt := TagType{true, "Apple", 2}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

}
