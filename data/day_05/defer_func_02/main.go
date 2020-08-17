package main

import (
	"fmt"
)

func demo() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func demo1() (x int) {
	defer func() {
		x++
	}()
	return 5
}
func demo2() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func demo3() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(demo())
	fmt.Println(demo1())
	fmt.Println(demo2())
	fmt.Println(demo3())
}
