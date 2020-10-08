package main

import (
	"flag"
	"fmt"
	"os"
)

func demo() {
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")

	// 命令行解析
	flag.Parse()
	fmt.Println(*name, *age, *married)
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, `testflag version: testflag/1.0.0
Usage: testflag [-h] [-p perfix]
Options:`)
		flag.PrintDefaults()
	}()
	if *name == "张三" && *age == 18 && *married == false && *name == "" && *age == 0 && *married == false {
		fmt.Println("请重新查看帮助信息")
	} else {
		demo1(*name, *age, *married)
	}

}
func demo1(name string, age int, married bool) {
	aa := name
	bb := age
	cc := married
	fmt.Printf("%s今年%d岁,婚姻状态为:%v", aa, bb, cc)
}
func main() {
	demo()
}
