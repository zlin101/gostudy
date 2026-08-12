//go:build ignore

package main

import "fmt"

// 结构体的定义
type Person struct {
	Name string
	Age  int
}
// 结构体的传递是值传递，传递的是副本，修改副本不会影响原来的结构体
func testPerson(p Person) {
	p.Name = "李四"
	p.Age = 30
}

// 结构体的指针传递，传递的是地址，修改会影响原来的结构体
func testPersonPointer(p *Person) {
	p.Name = "王五"
	p.Age = 40
}

func main() {
	// 第一种定义方式
	var p1 Person
	p1.Name = "张三"
	p1.Age = 20
	fmt.Println("p1=", p1)

	// 第二种定义方式
	p2 := Person{
		Name: "李四",
		Age:  30,
	}
	fmt.Println("p2=", p2)

	// 第三种定义方式
	p3 := new(Person)
	p3.Name = "王五"
	p3.Age = 40
	fmt.Println("p3=", *p3)

	// 第四种定义方式
	p4 := &Person{
		Name: "赵六",
		Age:  50,
	}
	fmt.Println("p4=", *p4)

	// 结构体的传递是值传递，传递的是副本，修改副本不会影响原来的结构体
	p5 := Person{
		Name: "孙七",
		Age:  60,
	}
	fmt.Println("p5=", p5)
	testPerson(p5)
	fmt.Println("p5=", p5)

	// 结构体的指针传递，传递的是地址，修改会影响原来的结构体
	p6 := Person{
		Name: "周八",
		Age:  70,
	}
	fmt.Println("p6=", p6)
	testPersonPointer(&p6)
	fmt.Println("p6=", p6)
}