package main

import (
	"fmt"
)

// slice的传递是引用传递，slice的底层数组是共享的
// 而且是动态的，slice的长度和容量是可以变化的
func testSlice(s []int) {
	s[0] = 100
}

func mainSlice() {
	s := []int{1, 2, 3}
	testSlice(s)
	fmt.Println(s)
}
