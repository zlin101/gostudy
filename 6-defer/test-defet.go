package main

import (
	"fmt"
)

// 验证return和defer的执行顺序, 以及defer的执行顺序是先进后出；
func testReturnDefer() int {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("return 1")
	return 1
}

func main() {
	fmt.Println(testReturnDefer())
}