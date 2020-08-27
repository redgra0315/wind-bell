package main

import "fmt"

/*
接口示例2
不管是什么牌子的车，都能跑
*/

// 定义一个car 接口类型
// 不管是什么结构体，只要有run 方法都能是car 类型
type car interface {
	run()
}
type fall struct {
	brand string
}

func (f fall) run() {
	fmt.Printf("%s 速度：70\n", f.brand)
}

type bailie struct {
	brand string
}

func (b bailie) run() {
	fmt.Printf("%s 速度：80\n", b.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	f1 := fall{
		brand: "法拉利",
	}
	p1 := bailie{
		brand: "保时捷",
	}
	drive(f1)
	drive(p1)
}
