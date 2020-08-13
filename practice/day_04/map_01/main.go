package main

import (
	"fmt"
	"strings"
)

func demo() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
func demo1() {
	m1 := make(map[string]int, 10)
	string := "how do you do"
	str := strings.Split(string, " ")

	for _, w := range str {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
func main() {
	demo1()
}
