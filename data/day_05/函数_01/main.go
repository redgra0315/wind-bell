package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 函数

// 无参无返回值
func demo() {
	fmt.Println("hello!")
}

//  有参数无返回值

func demo1(name string) {
	fmt.Println("hello", name)
}

// 有参数有返回值

func demo2(x, y int) int {
	a := x * y
	return a
}

// 可变参数
func demo3(title string, y ...int) int {
	fmt.Println(y)
	return 1
}

func demo4(s []int) []int {
	// 产生随机数
	rand.Seed(time.Now().Unix())
	for i := 0; i < 20; i++ {
		result := rand.Intn(100)
		s = append(s, result)
	}
	// 进行排序
	sort.Ints(s)
	return s
	//return 1
}

//  多个返回值
//func demo5()  {
//
//}
func main() {
	demo()
	demo1("杭州")
	a := demo2(4, 65)
	fmt.Println(a)
	//a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	a3 := make([]int, 0, 20)
	demo3("杭州", 1, 2, 3, 4, 5)
	aa := demo4(a3)
	fmt.Println(aa)
}
