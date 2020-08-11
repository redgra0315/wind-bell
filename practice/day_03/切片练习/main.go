package main

import (
	"fmt"
	"sort"
)

// 切片练习
func demo() {
	a := make([]int, 5, 10) // a  里面已经存在五个值的
	for i := 0; i < 10; i++ {
		a = append(a, i) // 这里是对a  进行append 添加元素
	}
	fmt.Println(a) // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
}

//数组进行排序
func demo1() {
	a := [...]int{3, 4, 7, 6, 12, 87, 8, 2, 0, 1}
	sort.Ints(a[:])
	fmt.Println(a)
}

func main() {
	demo()
	demo1()
}
