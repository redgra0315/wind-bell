package main

import "fmt"

// 函数全局作用域

// 定义全局变量
var x = 100

func demo() {
	// 函数中查找变量的顺序
	// 1、 先在函数内部查找，如果找到的话，就使用内部的
	// 2、如果在内部找不到就往外层进行查找
	// 3、 如果在全局也找不到，代码就报错
	//x := 10
	fmt.Println(x) // 输出100
}

// 局部变量
func demo1() {
	number := 100
	fmt.Println(number)
}

// 语句块作用域

func demo2() {
	if i := 10; i < 18 {
		fmt.Println("好好好")
	}
	fmt.Println(i)
}
func main() {
	demo()
	//fmt.Println(nmber)
	demo2()
}
