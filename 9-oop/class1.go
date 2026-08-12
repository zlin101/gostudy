//go:build ignore

package main

import "fmt"

// 类的定义
type Person struct {
	Name string
	Age  int
}
// 类的方法定义
func (this Person) GetName() string {
	fmt.Println("this=", this)
	return this.Name
}
// 类的指针方法定义
func (this *Person) GetAge() int {
	fmt.Println("this=", this)
	return this.Age
}	

// 类的传递是值传递，传递的是副本(副本：原对象的一个拷贝)，修改副本不会影响原来的类
func testPerson(this Person) {
	this.Name = "李四"
	this.Age = 30
}

// 类的指针传递，传递的是地址，修改会影响原来的类
func testPersonPointer(this *Person) {
	this.Name = "王五"
	this.Age = 40
}

func main() {
	p1 := Person{
		Name: "张三",
		Age:  20,
	}
	fmt.Println("p1=", p1)
	name := p1.GetName()
	age := p1.GetAge()
	fmt.Println("name=", name)
	fmt.Println("age=", age)
	
	testPerson(p1)
	fmt.Println("p1=", p1)
	
	testPersonPointer(&p1)
	fmt.Println("p1=", p1)	

}
