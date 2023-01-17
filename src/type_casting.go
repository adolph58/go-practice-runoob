package main

import "fmt"

// 类型转换用于将一种数据类型转换为另外一种类型的变量
func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为： %f\n", mean)
}
