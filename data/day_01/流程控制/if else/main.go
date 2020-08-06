package main

import "fmt"

//  if 条件判断
func demo() {
	age := 19
	if age >= 12 {
		fmt.Print("成年了")
	} else {
		fmt.Print("回家")
	}
}

func demo1() {
	age := 32
	if age >= 60 {
		fmt.Println("在家养老")
	} else if age >= 30 {
		fmt.Println("正在中年")
	} else if age >= 25 {
		fmt.Println("刚开始工作")
	} else {
		fmt.Println("好好上学")
	}
}
func demo2() {
	for {
		fmt.Print("1")
	}
}
func demo3() {
	s := "hello 沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
func main() {
	//demo()
	demo3()
}
