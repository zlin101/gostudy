//go:build ignore

package main

import "fmt"

// 接口的定义, 本质上是一个抽象类，里面定义了方法的集合，接口不能被实例化(指针)
type Animal interface {
	Eat()
	Move()
}

// Dog类的定义
type Dog struct {
	Name string
}

// Dog类实现了Animal接口的Eat方法
func (this *Dog) Eat() {
	fmt.Println(this.Name, "dog is eating")
}

// Dog类实现了Animal接口的Move方法
func (this *Dog) Move() {
	fmt.Println(this.Name, "dog is moving")
}

// Cat类的定义
type Cat struct {
	Name string
}

// Cat类实现了Animal接口的Eat方法
func (this *Cat) Eat() {
	fmt.Println(this.Name, "cat is eating")
}

// Cat类实现了Animal接口的Move方法
func (this *Cat) Move() {
	fmt.Println(this.Name, "cat is moving")
}

func main() {
	var a Animal

	//d := &Dog{Name: "旺财"}
	var d Dog
	d.Name = "旺财"	
	a = &d // 将Dog类型赋值给Animal接口变量a
	a.Eat()
	a.Move()

	//c := &Cat{Name: "咪咪"}
	var c Cat
	c.Name = "咪咪"
	a = &c // 将Cat类型赋值给Animal接口变量a
	a.Eat()
	a.Move()
}	

