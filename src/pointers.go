package main

import "fmt"

func main() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a        /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是：%x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量存储的指针地址：%x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值：%d\n", *ip)


	str := "Golang"
	var p *string = &str // p 是指向 str 的指针
	*p = "Hello"
	fmt.Println(str) // Hello 修改了 p，str 的值也发生了改变
}
