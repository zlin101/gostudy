//go:build ignore

// 对reflect包的使用

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (this Person) GetName() string {
	fmt.Println("this=", this)
	return this.Name
}

func (this *Person) GetAge() int {
	fmt.Println("this=", this)
	return this.Age
}

func ReflectTrace(arg interface{}) {
	t := reflect.TypeOf(arg)
	fmt.Println("Type:", t)

	v := reflect.ValueOf(arg)
	fmt.Println("Value:", v)
}

func main() {
	var num float64 = 3.1415926
	ReflectTrace(num)

	var p Person
	p.Name = "Alice"
	p.Age = 30
	ReflectTrace(p)

	t := reflect.TypeOf(p)
	fmt.Println("Type:", t)
	// 通过反射获取结构体的字段信息¥
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: %s %s\n", i, field.Name, field.Type)
	}
	// 通过反射获取结构体的方法信息
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("Method %d: %s %s\n", i, method.Name, method.Type)
	}

}
