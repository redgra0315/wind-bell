package main

import "fmt"

// iota 是够语言的常量计数器，只能在常量的表达式中使用
// iota 在const关键字出现时将被重置为0，const中每新增一行常量声明使iota 计数一次,使用iota能简化定义，在定义枚举时很有用。

// iota：
const (
	a1 = iota
	a2
	a3
	a4
)

const (
	b1 = iota // 0
	b2        // 1
	_         // 使用了匿名常量
	b4        // 3
)

const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4
)

const (
	d1, d2 = iota + 1, iota + 2 // 1,2
	d3, d4 = iota + 1, iota + 2 // 2,3
)

func main() {
	fmt.Println(a1)
	fmt.Println(a4)

	fmt.Println(b1, b2, b4)
	fmt.Println(c1, c2, c3, c4)
}
