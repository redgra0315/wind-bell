package main

import "fmt"

//使用make()函数构造切片

func demo() {
	a1 := make([]int, 10, 30)
	fmt.Printf("a1 := %d, len:%d,cap:%d\n", a1, len(a1), cap(a1))

	a2 := make([]int, 0, 30)
	fmt.Printf("a1 := %d, len:%d,cap:%d\n", a2, len(a2), cap(a2))
}

// 切片赋值
func demo1() {
	s3 := []int{1, 2, 3}
	s4 := s3
	fmt.Println(s3, s4)
	// s3 和s4 都指向了同一个底层数组，切片不保存值

	s3[0] = 123
	fmt.Println(s3, s4)
}

// 切片的遍历
func demo2() {
	s3 := []int{1, 2, 3, 4, 5, 67, 8, 9, 7767, 67, 45, 34}
	for i, v := range s3 {
		fmt.Println(i, v)
	}
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
}
func main() {
	demo2()
}
