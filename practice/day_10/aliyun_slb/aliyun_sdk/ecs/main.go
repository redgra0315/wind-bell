package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
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

func aliAuth() (client *sdk.Client) {
	//config := sdk.NewConfig().
	//	WithEnableAsync(true).
	//	WithGoRoutinePoolSize(5).
	//	WithMaxTaskQueueSize(1000)
	//credential := &credentials.BaseCredential{
	//	AccessKeyId:     "",
	//	AccessKeySecret: "",
	//}
	//
	//client, err := sdk.NewClientWithOptions("cn-hangzhou", config, credential)
	//if err !=nil{
	//	panic(err)
	//}

	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", "", "")
	if err != nil {
		panic(err)
	}
	return client
}
func getIp() {
	client := aliAuth()
	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Product = "Ecs"
	request.Domain = "ecs.aliyuncs.com"
	//request.ApiName = "DescribeInstances"
	request.ApiName = "DescribeInstanceStatus"
	request.QueryParams["PageNumber"] = "1"
	request.QueryParams["PageSize"] = "30"
	//request.QueryParams["RegionId"] = "cn-hangzhou"
	//request.QueryParams["RegionId"] = "cn-hangzhou"
	//request.TransToAcsRequest()
	//request.Scheme = "HTTPS"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())

}

func main() {
	//aliAuth()
	getIp()
}
