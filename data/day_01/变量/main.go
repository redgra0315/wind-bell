package main

import "fmt"

// Go 语言中推介使用驼峰命名
// 声明变量
//var names string
//var ages	int
//var notOk bool

// 批量全局声明
var (
	name string
	age  int
	isOk bool
)

// 声明局部变量同时赋值
func demo() {
	// 类型推导
	var s2 = "黄鹤楼"
	fmt.Println(s2)
	// 短变量声明
	s3 := "长江"
	fmt.Println(s3)
}

// 匿名变量

func main() {
	name = "煊赫门"
	age = 16
	isOk = true
	// Go 语言中变量声明了必须使用
	fmt.Println(name, age, isOk)
}
