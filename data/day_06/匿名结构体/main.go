package main

import "fmt"

// 嵌套结构体

type address struct {
	province string
	city     string
}
type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体
}

type company struct {
	name     string
	province string
	city     string
}

// 定义一个匿名结构体

func main() {
	var s struct {
		name string
		age  int
	}
	s.name = "北京"
	s.age = 1990
	fmt.Printf("%T %v\n", s, s)
	p1 := person{
		name: "周琳",
		age:  900,
		address: address{
			province: "杭州",
			city:     "萧山",
		},
	}
	fmt.Println(p1)
	// 查找顺序：先在自己结构体查找这个字段，找不到在到嵌套结构体中查找
	fmt.Println(p1.name, p1.city)
}
