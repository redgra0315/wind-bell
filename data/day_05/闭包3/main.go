package main

import (
	"fmt"
	"strings"
)

func make(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + "." + suffix
		}
		return name
	}
}
func main() {
	jpg := make("jpg")
	//test := make("bash.jpg")
	fmt.Println(jpg("test"))
	fmt.Println(jpg("tes12.jpg"))
}
