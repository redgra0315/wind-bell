package main

import "fmt"

// 使用值接受者和指针接收者实现接口的区别
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 方法使用值接受者
//func (c cat) move() {
//	fmt.Println("走猫步！")
//}
//func (c cat) eat(food string) {
//	fmt.Printf("猫吃:\n", food)
//}

// 使用指针接收者现实这个接口

func (c *cat) move() {
	fmt.Println("走猫步!")
}
func (c *cat) eat(food string) {
	fmt.Printf("猫吃:%s\n", food)
}
func main() {
	var a1 animal
	c1 := cat{
		name: "tom",
		feet: 4,
	}
	c2 := &cat{
		"假老练",
		4,
	}
	//a1 = c1 // 使用值类型
	a1 = &c1 // 实现animal这个接口的是cat的指针类型，
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
