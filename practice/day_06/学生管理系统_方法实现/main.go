package main

import (
	"fmt"
	"os"
)

type person struct {
	id   int8
	name string
	age  int8
}

func newPerson(id int8, name string, age int8) *person {
	return &person{
		id:   id,
		name: name,
		age:  age,
	}
}

var personMap = make(map[int8]*person, 50)

// 查看数据
func (p *person) showPerson() {
	fmt.Println("查看系统")
	for k, v := range personMap {
		fmt.Println(k, v.name, v.age)
	}
}

func (p *person) addPerson() {
	fmt.Println("Add person！")
	var (
		id   int8
		name string
		age  int8
	)
	fmt.Print("请输入ID:")
	fmt.Scan(&id)
	fmt.Print("请输入Name:")
	fmt.Scan(&name)
	fmt.Print("请输入age:")
	fmt.Scan(&age)
	per := newPerson(id, name, age)
	personMap[id] = per
	fmt.Print(*per)
}

func (p *person) delPerson() {
	fmt.Println("删除学生！")
	delID := int8(0)
	fmt.Print("输入删除的ID:")
	fmt.Scan(&delID)
	_, ok := personMap[delID]
	if !ok {
		fmt.Println("用户不存在！")
	} else {
		delete(personMap, delID)
		value, ok := personMap[delID]
		if !ok {
			fmt.Println("删除成功")
		} else {
			fmt.Println(value)
		}
	}
}

func main() {
	id := int8(0)
	fmt.Print("  ~~学生系统~~    ")
	for {
		fmt.Print(`
	1、查看
	2、添加
	3、删除
	4、退出
`)
		fmt.Print("你的选择是:")
		fmt.Scan(&id)
		switch id {
		case 1:
			show := person{}
			show.showPerson()
		case 2:
			add := person{}
			add.addPerson()
		case 3:
			del := person{}
			del.delPerson()
		case 4:
			os.Exit(0)
		}
	}
}
