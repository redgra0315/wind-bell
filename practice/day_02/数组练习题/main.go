package main

import (
	"fmt"
)

// 求数组元素的和

func demo() {
	a1 := [...]int{1, 3, 5, 7, 8}
	a2 := 0
	for _, v := range a1 {
		a2 = a2 + v
	}
	fmt.Println("数组元素的何为:", a2)

}

func demo1() {
	a1 := [...]int{1, 3, 5, 7, 8}

	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
func main() {
	//demo1()
	demo1()
}
