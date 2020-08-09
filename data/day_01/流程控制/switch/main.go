package main

import (
	"fmt"
)

// switch 简化大量的判断
func demo() {
	switch n := 10; {
	case n > 1:
		fmt.Print(n)
	case n < 3:
		fmt.Print(n)
	case n > 32:
		return
	case n >= 32:
		fmt.Print(n)
	default:
		fmt.Println(n)
	}
}
func main() {
	demo()
}
