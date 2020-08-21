package main

import "fmt"

// 结构体占用一块连续的内存
func demo3() {
	type x struct {
		a int8
		b int8
		c int8
		d string
	}
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
		d: string("hello"),
	}
	fmt.Printf("%p\n", &m.a)
	fmt.Printf("%p\n", &m.b)
	fmt.Printf("%p\n", &m.c)
	fmt.Printf("%p\n", &m.d)
}
func main() {
	demo3()
}
