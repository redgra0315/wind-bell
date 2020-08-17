package main

import "fmt"

// defer 函数

func demo() {
	fmt.Println("start")
	// defer 多用与函数结束之前释放资源(文件句柄、数据库连接、socket连接)
	defer fmt.Println("呵呵")
	fmt.Println("stop")
}

func main() {
	demo()
}
