package main

import "fmt"

// 引出接口的实例

type motion interface {
	speak() // 只要实现了speak 方法的变量都是speak类型
}
type cat struct{}

type dog struct{}

type person struct {
}

func (c cat) speak() {
	fmt.Println("喵喵!")
}
func (d dog) speak() {
	fmt.Println("旺旺！")
}

func (p person) speak() {
	fmt.Println("上天！")
}
func da(a motion) {
	// 接收一个参数，传进来什么，就操作什么
	a.speak()
}
func main() {
	//var c1 cat
	//var d1 dog
	//var p1 person
	c1 := cat{}
	d1 := dog{}
	p1 := person{}
	da(c1)
	da(d1)
	da(p1)

}
