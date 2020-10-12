package main

import (
	"flag"
	"fmt"
	"os"
)

func help() {
	// 定义帮助文档的信息
	var (
		serviceIp  string
		slbGroupId string
		port       int
		weight     int
		example    string
	)

	flag.StringVar(&serviceIp, "serviceIp", "172.16.0.0", "提供ecs服务器IP地址")
	flag.StringVar(&slbGroupId, "slbGroupId", "rsp-bp15k0sfl252m", "查看负载均衡groupId编号")
	flag.IntVar(&port, "port", 80, "请提供要修改服务的端口")
	flag.IntVar(&weight, "weight", 100, "权重小于99为服务下线，权重大于99为服务上线")
	flag.StringVar(&example, "example", "", "UpDateWeight serverId=172.16.0.0 port=80 weight=100")

	flag.Parse()
	fmt.Println(serviceIp, slbGroupId, port, weight, example)
}
func demo() {

	fmt.Println("In pogo  Project")

	var flagInt = flag.Int("flagname", 100, "help message  for flagname")
	var flagString = flag.String("flagstring", "name", "输入你的名字")

	var flagValue int
	flag.IntVar(&flagValue, "flagValue", 123, "In massage flagValue")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `pogo  version: pogo/1.0.0
Usage: pogo [-h] [-p Prefix]
Options: 
`)
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println(*flagInt, *flagString, flagValue)
	//fmt.Println(flagValue)
	//fmt.Println(flag.Arg(0))
	//fmt.Println(flag.Args())
	if flag.NArg() == 0 {
		fmt.Println("请查了帮助信息")
	} else {
		fmt.Println(flag.Arg(0))
		fmt.Println(flag.Args())
	}
}

func main() {
	//demo()
	help()
}
