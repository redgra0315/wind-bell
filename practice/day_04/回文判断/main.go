package main

import "fmt"

// 字符串从左往右和从右往左读是一样的，
func main() {
	ss := "上海自来水来自海上"
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}

	fmt.Println(r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回味")
}
