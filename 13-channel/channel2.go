//go:build ignore

// 有缓存的channel
package main

import (
	"fmt"
	"time"
)

func main (){
	defer fmt.Println("main is ended.")

	// 构建一个chan，cap=4，cap就是缓存的体现
	c := make(chan int, 4)
	fmt.Println("len = ", len(c), ", cap = ", cap(c))

	// 创建一个goroutine
	go func()  {
		defer fmt.Println("goroutine is ended.")
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("the ",i, "is sending to the chan;", "len | cap are", len(c), cap(c))
		}
	}()
	time.Sleep(2 * time.Second)
	for i:=0; i < 10; i++{
		num := <-c // 从c中取出数据
		fmt.Println("num = ", num)
	}

}
 