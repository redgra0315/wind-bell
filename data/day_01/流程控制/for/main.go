package main

import "fmt"

func demo2() {
	for {
		fmt.Print("1")
	}
}
func demo3() {
	s := "hello 沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}

// 跳出循环
func demo() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

// 跳过循环
func demo4() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}

func main() {
	demo4()
}
