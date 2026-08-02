package main

import (
	"fmt"
	
)

func main () {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(slice[1:3])
	fmt.Println(slice[1:])
	fmt.Println(slice[:3])
}
