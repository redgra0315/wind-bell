package main

import "fmt"

// 整型

func main() {
	// 十进制
	i1 := 101
	fmt.Printf("%b\n", i1) // 二进制
	fmt.Printf("%d\n", i1) // 十进制
	fmt.Printf("%x\n", i1) // 十进制转八进制
	fmt.Printf("%o\n", i1) // 十进制转十六进制

	// 八进制
	i2 := 071
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0xfff
	fmt.Printf("%d\n", i3)

	// 声明一个int 8
	i4 := int8(10)
	fmt.Printf("%T %d\n", i4, i4) // %T 输出类型
}
