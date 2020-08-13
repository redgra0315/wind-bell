package main

import "fmt"

// 匿名函数

var f1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	f1(10, 20)
	// 如果只是调用一次的函数，还可以简写位立即执行的函数
	func() {
		fmt.Println("haha")
	}()
}
