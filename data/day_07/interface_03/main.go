package main

import (
	"fmt"
)

// 接口的实现

type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("激动")
}

func (c chicken) eat(food string) {
	fmt.Printf("%s: 有%d条腿\n", food, c.feet)
}

func (c cat) move() {
	fmt.Println("走猫步")
}
func (c cat) eat(food string) {
	fmt.Printf("%s: 吃猫粮\n", food)
}

func main() {
	var a1 animal // 定义一个接口类型的变量
	bc := cat{    // 定义一个cat 类型的变量bc
		name: "蓝猫",
		feet: 4,
	}
	bd := chicken{
		feet: 2,
	}
	a1 = bc
	a1.move()
	a1.eat("蓝猫")
	a2 := bd
	a2.eat("战斗机")
	a2.move()
	//fmt.Println(a2)

}
