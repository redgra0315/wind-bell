package main

import "fmt"

type person struct {
	name, gender string
}

// go 语言中函数传参永远是拷贝
func demo(x person) {
	x.name = "武则天"
	x.gender = "女"
}

func demo1(x *person) {
	x.name = "武则天" // 根据内存地址找到那个原变量，修改的就是原来的变量
	x.gender = "女" // 语法糖，自动根据指针找对应的变量
}

func demo2() {
	var p3 = person{
		name:   "原酸",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)
}

func main() {
	var p person
	p.name = "李世明"
	p.gender = "男"
	fmt.Println(p)
	demo(p)
	fmt.Println(p)
	demo1(&p)
	fmt.Println(p)
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Println(p2)
	demo2()
}
