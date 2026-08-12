//go:build ignore

package main

import "fmt"

// map的定义
// 第一种定义方式
var m1 map[string]int

// 第二种定义方式
var m2 = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
}

// 第三种定义方式
var m3 = make(map[string]int)

// map为指针传递；
func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}
}

func main() {
	if m1 == nil {
		fmt.Println("m1 is nil")
	} else {
		fmt.Println("m1 is not nil")
	}
	fmt.Println("------------------------------")
	if m2 == nil {
		fmt.Println("m2 is nil")
	} else {
		fmt.Println("m2 is not nil")
		printMap(m2)
	}
	fmt.Println("------------------------------")
	if m3 == nil {
		fmt.Println("m3 is nil")
	} else {
		fmt.Println("m3 is not nil")
		m3["four"] = 4
		m3["five"] = 5
		printMap(m3)
	}
}
