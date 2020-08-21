package main

import "fmt"

// 方法和接收者

type dog struct {
	name string
}
type person struct {
	name string
	age  int
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

/*
方法是作用于特定类型的函数
接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
*/

func (d dog) wang(age int) (string, int) {
	return d.name, age
}

// 使用值接收者：传拷贝进去
func (p person) guonian() {
	p.age++
}

// 使用指针接收者：传内存地址进去
func (p *person) guonian2() {
	p.age++
}
func main() {
	d1 := newDog("小白")
	hobby, age := d1.wang(12)
	fmt.Println(hobby, age)
	p1 := newPerson("原生态", 1980)
	p1.guonian()
	fmt.Println(p1.age)
	p1.guonian2() // 相当于 &p1.gounian2()
	fmt.Println(p1.age)
}
