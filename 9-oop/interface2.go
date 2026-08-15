
//go:build ignore

package main

import "fmt"

// 测试空接口类型
func Test (arg interface{}) {
	fmt.Println("arg=", arg)
	value, ok := arg.(int)
	if ok {
		fmt.Println("arg is int, value=", value)
	} else {
		fmt.Println("arg is not int")
	}
}

type Book struct {
	Title  string
} 

func main() {
	Test(100)
	Test("Hello")
	Test(3.14)
	Test(true)
	Test([]int{1, 2, 3})
	Test(9)

	b := Book{Title:  "Go语言编程"}
	Test(b)
}