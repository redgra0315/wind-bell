package main

import (
	"fmt"
)

/*
数组

数组是值类型
存放元素的容器
必须指定存放元素的类型合容量(长度)
数组的长度是数组类型的一部分
多维数组只能存放相同类型
*/

// 数组的定义
func demo() {
	var a1 [3]bool // [true  false true]
	var a2 [4]bool // [true true false  false]

	fmt.Printf("%T %T\n", a1, a2)
}

// 数组初始化
// bool 值默认位：false

func demo1() {
	// 第一种初始化
	a1 := [3]bool{true, true, false}

	// 第二种初始化
	// ...  根据初始值自动推断数组的长度是多少
	a2 := [...]int{1, 3, 4, 5, 6, 7, 8, 9, 45, 2, 223}

	// 第三中初始化
	// 根据索引来进行初始化
	a3 := [5]int{0: 1, 3: 2}

	// 打印
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}

// 数组的遍历

func demo2() {
	city := [...]string{"北京", "上海", "深圳"}
	// 根据索引来进行遍历
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}

	// range 遍历
	for i, v := range city {
		fmt.Println(i, v)
	}
}

// 多维数组
func demo3() {
	a1 := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	//fmt.Println(a1)

	// 多维数组的遍历
	for _, v := range a1 {
		fmt.Println(v)
		for _, v1 := range v {
			fmt.Println(v1)
		}
		//fmt.Println(v)
	}

}

// 数组是值类型
func demo4() {
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[1] = 100
	fmt.Println(b1)
	fmt.Println(b2)
}
func main() {
	//demo()
	//demo1()
	//demo2()
	//demo3()
	demo4()
}
