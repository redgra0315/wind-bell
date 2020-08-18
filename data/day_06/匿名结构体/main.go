package main

import "fmt"

// 定义一个匿名结构体

func main() {
	var s struct {
		name string
		age  int
	}
	s.name = "北京"
	s.age = 1990
	fmt.Printf("%T %v\n", s, s)
}
