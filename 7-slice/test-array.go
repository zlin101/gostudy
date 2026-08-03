//go:build ignore

package main

import "fmt"

// 不同长度的array，是不同的数据类型
func typeVs() {
	var a [3]int
	var b [4]int
	fmt.Printf("a type: %T, b type: %T\n", a, b)
}

// 数组的传递是值传递，传递的是数组的副本
func arrayPass() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println("before pass:", a)
	testArray(a)
	fmt.Println("after pass:", a)
}

func testArray(arr [3]int) {
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
}

func main() {
	typeVs()
	arrayPass()
}
