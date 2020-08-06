package main

import "fmt"

// 浮点数

func main() {
	//math.MaxFloat32
	f1 := 1.23455
	fmt.Printf("%T %v\n", f1, f1) // 默认Go语言中的小数都是float64 类型

	//定义float 32 位
	f2 := float32(1.23455)
	fmt.Printf("%T %v\n", f2, f2)
}
