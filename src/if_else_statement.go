package main

import "fmt"

func main() {
	/* 局部变量定义 */
	var a int = 100

	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为：%d\n", a)
}
