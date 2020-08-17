package main

import "fmt"

func demo() func(int, int) int {
	return func(x, y int) int {
		return x * y
	}
}

func demo1(a, b int) int {
	return a + b
}

func main() {
	ds := demo()
	fmt.Printf("%T %d\n", ds, ds)
}
