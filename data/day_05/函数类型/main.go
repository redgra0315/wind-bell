package main

import "fmt"

// 函数类型

func demo() {
	fmt.Println("hello 沙河")
}

func demo1() int {
	return 10
}

// 函数也可以作为参数的类型

func demo2(x func() int) {
	ret := demo1()
	fmt.Println(ret)
}

func demo3(x, y int) int {
	return x + y
}

// 函数还可以作为返回值
func demo4(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	a := demo
	fmt.Printf("%T\n", a)
	b := demo1
	fmt.Printf("%T\n", b)
	demo2(demo1)
	fmt.Printf("%T\n", demo3)
	c := demo4(demo1)
	fmt.Printf("%T\n", c)
}
