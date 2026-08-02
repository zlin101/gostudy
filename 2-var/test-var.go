package main

import "fmt"

func main() {
	fmt.Println("this is a var test!")

	// 第一种变量声明方式，默认是0
	var a int
	fmt.Println(a)

	// 第二种变量声明方式
	var b int = 20
	fmt.Println(b)

	// 第三种变量声明方式
	var c = 30
	fmt.Println(c)
	fmt.Printf("c的类型是%T\n", c)

	// 第四种变量声明方式, 这种方式只能在函数体中使用，但是很常用
	d := 40
	fmt.Println(d)
	fmt.Printf("d的类型是%T\n", d)

	// 声明多个变量
	var e, f, g int = 50, 60, 70
	fmt.Println(e, f, g)

	// 声明多个变量，类型相同，可以省略类型
	var h, i, j = 80, 90, 100
	fmt.Println(h, i, j)

	// 声明多个变量，类型不同，可以省略类型
	var k, l, m = 110, "hello", true
	fmt.Println(k, l, m)
	fmt.Printf("k的类型是%T, l的类型是%T, m的类型是%T\n", k, l, m)

	// 多行声明变量
	var (
		n int    = 120
		o string = "world"
		p bool   = false
	)
	fmt.Println(n, o, p)
	fmt.Printf("n的类型是%T, o的类型是%T, p的类型是%T\n", n, o, p)
}
