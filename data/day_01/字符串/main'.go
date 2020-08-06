package main

import (
	"fmt"
	"strings"
)

// 字符串操作
func demo1() {
	// 字符串拼接
	name := "杭州"
	world := "互联网"
	// 直接打印
	fmt.Println(name + world)

	// 进行拼接，有返回值
	worlds := fmt.Sprintf("%s%s\n", name, world)
	fmt.Println(worlds)

	// 分割
	aaa := "234234234234324"
	str := strings.Split(aaa, "")
	fmt.Println(str)
	fmt.Println(strings.HasPrefix(aaa, "2"))
	fmt.Println(strings.HasSuffix(aaa, "2"))
	fmt.Println(strings.Contains(aaa, "2"))
	fmt.Println(strings.ContainsRune(aaa, 2))
}

func main() {
	demo1()
}
