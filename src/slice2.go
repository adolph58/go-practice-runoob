package main

import "fmt"

func main() {
	// 创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	// 打印原始切片：0-8
	fmt.Println("numbers ==", numbers)

	// 打印子切片从索引1（包含）到索引4（不包含）：1-3
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	// 默认下限为 0：0-2
	fmt.Println("numbers[:3]", numbers[:3])

	// 默认上限为 len(s)：4-8
	fmt.Println("number[4:]", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	// 打印子切片从索引0（包含）到索引2（不包含）： 0-1
	numbers2 := numbers[:2]
	printSlice(numbers2)

	// 打印子切片从索引2（包含）到索引5（不包含）： 2-4
	numbers3 := numbers[2:5]
	printSlice(numbers3)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
