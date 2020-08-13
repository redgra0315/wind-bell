package main

import (
	"fmt"
)

// 闭包

func demo(f func()) {
	fmt.Println("this is: demo()")
	f()
}
func demo1(x, y int) {
	fmt.Println("this is demo1()")
	fmt.Println(x, y)
}

// 定义一个函数对demo1 包装
//func demo2(f func(x,y int)) func() {
//	tmp := func(){
//		f()
//		fmt.Println("hello")
//	}
//	return tmp
//}

func demo3(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}
func main() {
	ret := demo3(demo1, 101, 200)
	ret()
}
