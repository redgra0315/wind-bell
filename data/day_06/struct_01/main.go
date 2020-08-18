package main

import "fmt"

// 结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func demo() {
	// 声明一个persion类型的变量p
	var p person
	// 通过字段进行赋值
	p.name = "周琳"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	fmt.Println(p.name)
}
func main() {
	demo()
}
