package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func demo() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file:%s\n", err))
	}

	viper.WatchConfig()

	//aa  := viper.ReadInConfig()
	//fmt.Println(aa)
	//name1 := viper.Get("name")
	//fmt.Println(name1)
	fmt.Println(viper.GetStringMapString("person"))
	fmt.Println(viper.GetStringMapString("region.city"))
	//fmt.Println(viper.GetString("name"))
	//fmt.Println(viper.GetInt("age"))
	//fmt.Println(viper.GetString("hobby"))
}

func demo2() {
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	fmt.Println(viper.GetInt("flagname"))
}
func main() {
	demo()
}
