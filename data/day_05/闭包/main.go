package main

import (
	"fmt"
)

// 闭包

func aa() {
	fmt.Println("aa")
}
func demo(f func()) {
	fmt.Println("this is: demo()")
	f()
}
func demo1(x, y int) {
	fmt.Println("this is demo1()")
	//a = x
	//b = y
	//c = func() {
	fmt.Println(x + y)
	//}
	//return
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

func demo4(x func(func(), int, int), m, n int) (func(), int, int) {
	y1 := func() {
		fmt.Println("y1")
	}
	tmp := func() {
		x(y1, 100, 300)
	}
	return tmp, m, n
}
func main() {
	ret := demo3(demo1, 101, 200)
	ret()
	//ret1, ret2,ret5  := demo4(demo1,32,43)
	//fmt.Println(ret1,ret2,ret5)
	//ret3,ret4 := demo1(12,43)
	//fmt.Println(ret4,ret3)
}
