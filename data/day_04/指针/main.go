package main

import "fmt"

func demo() {
	// &:取地址
	n := 10
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)

	// *:根据地址取值
	m := *p
	fmt.Println(&m)
}

// 使用new 申请内存地址
func demo1() {
	a := new(int)
	*a = 100
	fmt.Println(*a)
}
func main() {
	demo1()
}
