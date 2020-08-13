package main

import "fmt"

// 元素类型为map 的切片1
func demo() {
	var s1 = make([]map[int]string, 10)
	s1[0] = make(map[int]string, 10)
	s1[1] = make(map[int]string, 10)
	s1[0][10] = "杭州"
	s1[1][11] = "苏州"
	fmt.Println(s1)
}

// 元素类型为map 的切片2
func demo2() {
	a := make(map[int]string)
	a[0351] = "哈能藏"
	a[010] = "东扽"

	a1 := make(map[int]string)
	a[03512] = "哈fd"
	a[0102] = "东扽"

	s1 := make([]map[int]string, 0, 10)
	//a3 := make([]map[int]string,10,20)
	//s1[0] = make(map[int]string)
	//s1[1] = make(map[int]string)
	//s1[100] = make(map[int]string)
	s1 = append(s1, a)
	//s1 = append(s1,a1)
	//s1[0][1] = "杭州"
	//s1[1][2] = "杭州"
	//s1[1][11] = "苏州"
	s1 = append(s1, a1)

	fmt.Println(s1)
}

// 值为切片类型的map
func demo3() {
	m1 := make(map[string][]int, 10)
	m1["北京"] = []int{12, 34, 34, 54}
	fmt.Println(m1)
}
func main() {
	demo2()
}
