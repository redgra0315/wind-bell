package main

import "fmt"

// 同一个结构体可以实现多个接口

// 接口嵌套
type animal interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat 实现了mover接口
func (c cat) mover() {
	fmt.Println("走猫步!")
}

// cat实现了eater接口
func (c cat) eater(food string) {
	fmt.Printf("猫吃:%s\n", food)
}

func main() {

}
