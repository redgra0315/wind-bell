package main

import "fmt"

// 数值类型的练习

func demo() {
	name := "张三"
	age := 18
	grade := 59.9
	maritalStatus := false
	fmt.Printf("name: %s  %T\n", name, name)
	fmt.Printf("age: %d  %T\n", age, age)
	fmt.Printf("grade: %v  %T\n", grade, grade)
	fmt.Printf("maritalStatus: %v  %T\n", maritalStatus, maritalStatus)
}
func main() {
	demo()
}
