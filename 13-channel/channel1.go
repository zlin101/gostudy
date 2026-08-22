//go:build ignore
// 无缓存
package main

import "fmt"

func main() {
	defer fmt.Println("the main is ended. ")
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine ended.")

		fmt.Println("goroutine is running...")

		c <- 678 // 将数据通过chnanel就行传输；
	}()
	num := <-c // 从c中接收数据，并赋值给num；注意，main & goroutine 会因为chnanel导致的互相阻塞；
	// 所以goroutine ended.一定发生在chnanel传输成功之后；
	// 为什么是互相阻塞？ main 和 goroutine 他们谁先running到chan传输的步骤都不一定，先到的必须等待后到的；

	fmt.Println("the num = ", num)

}
