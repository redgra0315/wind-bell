package main

import (
	"encoding/json"
	"fmt"
)

/*
结构体与json
1、序列化：把GO语言中的结构体变量 --> json格式的字符串
2、发序列化：把json格式的字符串  --> GO语言中能够识别的结构体变量
*/

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func demo() {
	p1 := person{
		Name: "周琳",
		Age:  900,
	}
	// 序列化与反序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	// 反序列化
	str := `{"name":"周琳1","age":90}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在json.Unmarshal内部修改p2的值

	fmt.Printf("%#v\n", p2)
}

func main() {
	demo()
}
