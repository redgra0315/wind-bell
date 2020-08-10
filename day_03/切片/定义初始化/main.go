package main

import (
	"fmt"
)

/*
切片的定义
声明切片类型的基本语法如下：
var name []T
*/

func demo() {
	// 定义一个存放int类型元素的切片
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)

	// 初始化
	s1 = []int{1, 2, 3, 4, 5}
	s2 = []string{"北京", "杭州", "上海"}

	fmt.Println(s1, s2)
}

// 切片的长度和容量
func demo1() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []string{"北京", "杭州", "上海"}
	fmt.Printf("len:%d cap:%d\n", len(s1), cap(s1))
	fmt.Printf("len:%d cap:%d\n", len(s2), cap(s2))
}

// 有数组得到切片

func demo2() {
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	a2 := a1[0:5]
	fmt.Println(a2)
}
func main() {
	demo2()

}
