package main

import "fmt"

// 九九乘法表
func demo() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf(" %d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
func main() {
	demo()
}
