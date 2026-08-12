//go:build ignore

package main

import "fmt"


// slice 的扩容和截取
// 第一种定义方式
var s1 []int

// 第二种定义方式，注意“:=”只可以在函数内部使用;
var s2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 第三种定义方式, 仅指定长度，未指定容量：
var s3 = make([]int, 5)

// 第四种定义方式，指定容量和长度：
var s4 = make([]int, 2, 3)

func main() {
	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not nil")
	}
	fmt.Println("------------------------------")
	if s2 == nil {
		fmt.Println("s2 is nil")
	} else {
		fmt.Println("s2 is not nil")
		for i := 0; i < len(s2); i++ {
			fmt.Println("value=", s2[i])
		}
	}	
	fmt.Println("------------------------------")
	if s3 == nil {
		fmt.Println("s3 is nil")
	} else {
		fmt.Println("s3 is not nil")
		for i := 0; i < len(s3); i++ {
			fmt.Println("value=", s3[i])
		}
	}
	fmt.Println("------------------------------")
	if s4 == nil {
		fmt.Println("s4 is nil")
	} else {
		fmt.Println("s4 is not nil")
		for i := 0; i < len(s4); i++ {
			fmt.Println("value=", s4[i])
		}
	}
	// 以s4为例，就行扩容和截取
	fmt.Println("------------------------------")
	fmt.Println("最开始s4的长度为：", len(s4))
	fmt.Println("最开始s4的容量为：", cap(s4))
	// 扩容(cap内容量为3，超过3就会扩容)
	s4 = append(s4, 1)
	fmt.Println("第一次追加s4的长度为：", len(s4))
	fmt.Println("第一次追加s4的容量为：", cap(s4))
	// 第二次扩容；
	s4 = append(s4, 2)
	fmt.Println("第二次追加s4的长度为：", len(s4))
	fmt.Println("第二次追加s4的容量为：", cap(s4))
	
	// 截取：[,) 左闭右开区间
	s4 = s4[0:5]
	fmt.Println("s4的长度为：", len(s4))
	fmt.Println("s4的容量为：", cap(s4))

	// 深浅拷贝
	// 浅拷贝：slice的底层数组是共享的，修改一个slice的值会影响另一个slice的值
	s5 := s4
	s5[0] = 100
	fmt.Println("s4的值为：", s4)
	fmt.Println("s5的值为：", s5)

	// 深拷贝：使用copy函数进行拷贝，两个slice的底层数组不共享，修改一个slice的值不会影响另一个slice的值
	s6 := make([]int, len(s4))
	copy(s6, s4)
	s6[0] = 200
	fmt.Println("s4的值为：", s4)
	fmt.Println("s6的值为：", s6)	
	
}

