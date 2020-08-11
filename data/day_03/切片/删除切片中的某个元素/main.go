package main

import "fmt"

func demo() {
	a1 := [...]int{1, 2, 3, 5, 5, 7, 3, 456, 4, 4, 23}
	a2 := a1[:]
	a2 = append(a2[0:1], a2[5:]...)
	fmt.Println(a2)
	fmt.Println(a1)
}
func main() {
	demo()
}
