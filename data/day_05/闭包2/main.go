package main

import "fmt"

// 闭包是什么
// 闭包是一个函数，这个函数包含了它外部作用域的一个变量
func adder(x int) func(int) int {
	//x := 100
	return func(i int) int {
		x += i
		return x
	}
}
func main() {
	ret := adder(21)
	fmt.Println(ret(100))
}
