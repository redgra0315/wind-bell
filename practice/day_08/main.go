package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	PageNumber int       `json:"PageNumber"`
	Instances  Instances `json:"Instances"`
}

type Instances struct {
	Instance []Instance `json:"instance"`
}

type Instance struct {
	VpcAttributes VpcAttributes `json:"VpcAttributes"`
}

type VpcAttributes struct {
	PrivateIpAddress PrivateIpAddress `json:"PrivateIpAddress"`
}

type PrivateIpAddress struct {
	IpAddress []string `json:"IpAddress"`
}

func main() {
	b, err := ioutil.ReadFile("practice\\day_08\\ecs.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	//method1(b)
	method2(b)
}

// 方法1
func method1(b []byte) {
	var data Data
	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data.Instances.Instance[0].VpcAttributes.PrivateIpAddress.IpAddress[0])
}

// 方法2
func method2(b []byte) {
	var data map[string]json.RawMessage
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Println(err)
		return
	}
	var Instances map[string]json.RawMessage
	err := json.Unmarshal(data["Instances"], &Instances)
	if err != nil {
		fmt.Println(err)
		return
	}

	var Instance []Instance
	err = json.Unmarshal(Instances["Instance"], &Instance)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Instance[0].VpcAttributes.PrivateIpAddress.IpAddress[0])
}
