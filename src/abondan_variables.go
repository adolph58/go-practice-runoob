package main

import "fmt"

/**
 * 用 "_" 舍弃值
 */

func main() {
	_, numb, strs := numbers() // 只获取函数返回值的后两个
	fmt.Println(numb, strs)
}

func numbers()(int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
