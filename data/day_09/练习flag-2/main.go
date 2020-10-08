package main

import (
	"flag"
	"fmt"
	"os"
)

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
	fmt.Println(flag.Arg(0))
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
}
func main() {
	demo()
}
