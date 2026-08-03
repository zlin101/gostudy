package main

import "fmt"

// slice 的传递是引用传递，底层数组共享
// 这里内联 testSlice，使 main.go 自包含、可独立编译运行
func testSlice(s []int) {
	s[0] = 100
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	testSlice(s)
	fmt.Println(s)
}
