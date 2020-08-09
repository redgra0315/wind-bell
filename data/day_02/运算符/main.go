package main

import "fmt"

// 运算符
func demo() {
	/*	位运算符，针对的是二进制数
		&： 按位与
		|: 按位或.两位有一个1就为1，
		^: 按位异或(两位不一样则为1)
		<<: 将二进制位左移指定位数
	*/
	fmt.Println(5 & 2)
	fmt.Println(5 | 2)
	fmt.Println(5 ^ 2)
	fmt.Println(5 << 1)
	fmt.Println(1 << 10) // 1024
	//二进制右移指定的位数
	fmt.Println(5 >> 1)
	var m = int8(1)
	fmt.Println(m << 5)

}
func main() {
	demo()
}
