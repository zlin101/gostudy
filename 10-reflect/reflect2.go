//go:build ignore

package main

import "fmt"

// 接口的定义
type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// Book类的定义
type Book struct {
	Title string
}

// Book类实现了Reader接口的ReadBook方法
func (this *Book) ReadBook() {
	fmt.Println(this.Title, "is being read")
}

// Book类实现了Writer接口的WriteBook方法
func (this *Book) WriteBook() {
	fmt.Println(this.Title, "is being written")
} 

func main() {
	// b: pair<type Book, value *Book>
	b := &Book{Title: "Go语言编程"}
	// r: pair<type Reader, value *Book>
	var r Reader // 定义一个Reader接口变量r
	// r: pair<type Book, value *Book>
	r = b // 将Book类型赋值给Reader接口变量r
	r.ReadBook()
	// w: pair<type Writer, val ue *Book>
	var w Writer // 定义一个Writer接口变量w
	// w: pair<type Book, value *Book>
	w = r.(Writer) // 将Book类型赋值给Writer接口变量w
	w.WriteBook()
	// 总结：
	// 1. 接口类型的变量可以存储实现了该接口的任意类型的值
	// 2. 接口类型的变量存储的是一个pair<type, value>，type是实现了该接口的类型，value是该类型的值
	// 3. 接口类型的变量可以调用该接口定义的方法，实际调用的是存储的值的类型的方法
	// 4. 接口类型的变量可以被赋值为实现了该接口的任意类型的值
}
