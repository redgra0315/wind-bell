package main

import "fmt"

// 常量
// 定义了常量之后不能修改
const version = "V0.0.1"

// 批量声明
const (
	statusOk = 200
	notFound = 404
)

//  批量声明，如果某一行声明后没有写值，默认和上一个值相同
const (
	n1 = 100
	n2
	n3
)

func main() {
	fmt.Println(version)
	fmt.Println(statusOk)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}
