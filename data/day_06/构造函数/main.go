package main

import "fmt"

// 构造函数：返回一个结构体变量的函数
// 结构体属于值类型： 赋值的时候相当于拷贝

// 如果结构体中的参数较少的情况下可以使用返回结构体
// 如果结构体中的参数较多的情况下，建议返回结构体指针
type person struct {
	name string
	age  int
}

/*
构造函数： 约定成俗用new开头
返回的是结构体还是结构体指针
当结构体比较大的时候尽量使用结构体指针，减少程序的运行内存开销
*/
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func main() {
	p1 := newPerson("杭州", 0531)
	p2 := newPerson("上海", 012)
	fmt.Println(p1)
	fmt.Println(p2)
}
