package main

import "fmt"

// 修改字符串
func demo1() {
	s1 := "big"
	// 把字符串强制转换成一个rune切片
	s2 := []rune(s1)
	s2[0] = 'a' // 字符
	// 把rune 切片强制转换成字符串
	fmt.Println(string(s2))

	c1 := "红" // c1:string
	c2 := '红' // c2:int32 (rune)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
}

// 类型转换
func demo2() {
	d1 := 12
	var d2 float64
	d2 = float64(d1)
	fmt.Printf("%T\n", d2)
}
func main() {
	//demo1()
	demo2()
}
