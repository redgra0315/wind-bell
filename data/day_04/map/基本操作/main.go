package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map
func demo() {
	// map 必须进行初始化才能使用
	m1 := map[string]int{}

	m1["区号"] = 0351
	m1["年龄"] = 1872
	fmt.Println(m1)
}

// 使用make 进行初始化

func demo1() {
	var a map[string]int
	a = make(map[string]int, 10)
	a["杭州"] = 0351
	a["北京"] = 010
	fmt.Println(a) //  如果key不存在，返回对应类型的空值
	// 约定成俗用ok接收返回的布尔值
	value, ok := a["杭州"]
	if !ok {
		fmt.Println("不存在")
	} else {
		fmt.Println(value)
	}

	//	map 遍历
	for i, v := range a {
		fmt.Println(i, v)
	}

	//	使用delete()函数删除键值对
	delete(a, "杭州")
	fmt.Println(a)
	//	如果删除map 中不存在的key,delete 不做操作
}

// 遍历map 产生随机数
func demo2() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	var scoreMap = make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	//fmt.Println(scoreMap)

	//取出map 中所有的key存入切片keys
	keys := make([]string, 0, 100)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//fmt.Println(keys)
	//	对切片进行排序
	sort.Strings(keys)
	//fmt.Println(keys)
	for _, key := range keys {

		fmt.Println(key, scoreMap[key])
	}
}
func main() {
	demo2()
}
