package main

import "fmt"

// 结构体嵌套

type addr struct {
	province string
	city     string
}
type student struct {
	name string
	addr // 匿名嵌套别的结构体
	//addr	addr     嵌套别的结构体
}

func main() {
	p := student{
		name: "tom",
		addr: addr{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Print(p)
}
