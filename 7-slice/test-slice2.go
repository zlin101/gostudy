//go:build ignore

package main

import (
	"fmt"
)

// slice 的声明方式
func deSlice() {
	// 1. 声明一个空的 slice
	var s1 []int
	fmt.Println(s1) // []

	// 2. 声明一个有初始值的 slice
	s2 := []int{1, 2, 3}
	fmt.Println(s2) // [1 2 3]

	// 3. 使用 make 函数创建 slice
	s3 := make([]int, 5) // 创建一个长度为 5 的 slice，初始值为零值
	fmt.Println(s3)      // [0 0 0 0 0]

	// 4. 使用 make 函数创建 slice，并指定容量
	s4 := make([]int, 3, 5) // 创建一个长度为 3，容量为 5 的 slice
	fmt.Println(s4)         // [0 0 0]
}

func main() {
	deSlice()
	var s1 []int
	fmt.Println(s1) // []

	s2 := []int{}
	fmt.Println(s2) // []

	if s1 != nil {
		fmt.Println("s1 is not nil")
	} else {
		fmt.Println("s1 is nil")
	}  
	if s2 != nil {
		fmt.Println("s2 is not nil")
	} else {
		fmt.Println("s2 is nil")
	}

}
