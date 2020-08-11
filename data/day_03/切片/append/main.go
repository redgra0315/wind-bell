package main

import "fmt"

// 切片扩容 append()
func demo() {
	a1 := []string{"北京", "上海", "杭州"}
	fmt.Printf("a1=%v len:%d cap:%d\n", a1, len(a1), cap(a1))
	a1 = append(a1, "广州")
	fmt.Println(a1)
	fmt.Printf("a1=%v len:%d cap:%d\n", a1, len(a1), cap(a1))
}
func main() {
	demo()
}
