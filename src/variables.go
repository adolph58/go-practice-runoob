package main

import "fmt"

func main() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	/*
	 * 变量声明
	 */
	// 第一种，指定变量类型，如果没有初始化，则变量默认为零值

	// 声明一个变量并初始化
	var d = "RUNOOB"
	fmt.Println(d)

	// 没有初始化就为零值
	var e int
	fmt.Println(e)

	// bool 零值为 false
	var f bool
	fmt.Println(f)

	var g int
	var h float64
	var i bool
	var j string
	fmt.Printf("%v %v %v %q\n", g, h, i, j)

	// 第二种，根据值自行
	var k = true
	fmt.Println(k)

	// 第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生变量错误，格式：
	//var intVal int
	//intVal := 1

	// 相当于 var intVal int
	// intVal = 1
	intVal := 1
	fmt.Println(intVal)

	// 可以将 var f string = "Runoob" 简写为 f:= "Runoob"
	l := "Runoob" // var f string = "Runoob"
	fmt.Println(l)

}
