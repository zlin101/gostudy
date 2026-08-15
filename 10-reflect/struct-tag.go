//go:build ignore

package main

import (
	"fmt"
	"reflect"
)
// Tag 通常采用 key:"value" 的形式，key 是标签的名称，value 是标签的值。标签可以包含多个键值对，每个键值对之间用空格分隔。标签的值通常是一个字符串，可以包含任意字符，包括空格和特殊字符。
type Person struct {
	Name string `json:"name" docs:"姓名"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{
		Name: "张三",
		Age:  20,
	}
	
	t := reflect.TypeOf(p)
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: %s %s Tag: %s\n", i, field.Name, field.Type, field.Tag)
	}
}
