//go:build ignore

package main

import "fmt"

// 父类(nomal)的定义
type Human struct {
	Name string
	Age  int
}

func (this *Human) Eat() {
	fmt.Println(this.Name, "human have to eat")
}

func (this *Human) Walk() {
	fmt.Println(this.Name, "human can walk")
}
// 子类定义
type Student struct {
	Human  // 匿名字段，继承了Human的属性和方法(默认继承，区别C++的public、protected、private)
	School string
}

func (this *Student) Study() {
	fmt.Println(this.Name, "student have to study in", this.School)
}

func main() {
	/*
	s1 := Student{
		Human: Human{
			Name: "张三",
			Age:  20,
		},
		School: "清华大学",
	}
	*/
	// 第二种定义方式
	var s1 Student
	s1.Name = "张三"
	s1.Age = 20
	s1.School = "清华大学"

	fmt.Println("s1=", s1)
	s1.Eat()
	s1.Walk()
	s1.Study()
}