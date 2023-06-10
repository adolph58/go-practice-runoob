package main

import "fmt"
/**
Go 语言中，并不需要显式地声明实现了哪一个接口，只需要直接实现该接口对应的方法即可
 */

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}
	stu := p.(*Student) // 接口转为实例
	fmt.Println(p.getName()) // Tom
	fmt.Println(stu.age)
}