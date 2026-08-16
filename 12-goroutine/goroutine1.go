//go:build ignore

package main

import (
	"fmt"
	"runtime"
	"time"
)

// 测试runtime.Goexit的使用
func main (){
	go func() {
		defer fmt.Println("A.defer")
		go func ()  {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B")
		}()
		runtime.Goexit()
		fmt.Println("A")
	}()
	time.Sleep(1 * time.Second)  // 等 goroutine 执行完
	fmt.Println("main end")

}