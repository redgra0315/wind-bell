package main

import "fmt"

// 给自定义类型加方法
// 不能给别的包里面的类型添加方法，只能给自己包里的类型添加方法
type Myint int

func (m Myint) hello() {
	fmt.Println("哈哈哈", m)
}
func main() {
	m := Myint(100)
	m.hello()
}
