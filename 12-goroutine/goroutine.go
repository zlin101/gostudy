//go:build ignore

package main

import (
	"fmt"
	"time"
)

// 测试goroutine的基本使用
func main() {
	// 启动一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}() // 立即执行匿名函数

	// 主goroutine继续执行
	for i := 0; i < 5; i++ {
		fmt.Println("Main Goroutine:", i)
		time.Sleep(150 * time.Millisecond)
	}

	// 等待一段时间，确保goroutine执行完毕
	time.Sleep(1 * time.Second)
}