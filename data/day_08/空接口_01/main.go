package main

import "fmt"

// 定义一个空接口

//interface: 关键字
//interface{}:空接口类型
func demo() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "北京"
	m1["age"] = 18
	m1["grade"] = 87.3
	m1["status"] = true
	m1["hobby"] = [...]string{"说", "学", "逗", "唱"}
	fmt.Println(m1)
}

// 空接口作为函数参数

func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
func main() {
	demo()
	show("你好")
	show(3243)
	show(3.43)
	show(nil)
	show(demo)
}
