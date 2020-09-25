package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func main() {
	config := sdk.NewConfig().
		WithEnableAsync(true).
		WithGoRoutinePoolSize(5).
		WithMaxTaskQueueSize(1000)
	credential := &credentials.BaseCredential{
		AccessKeyId:     "",
		AccessKeySecret: "",
	}
	client, err := sdk.NewClientWithOptions("cn-hangzhou", config, credential)
	//client, err := sdk.NewClientWithAccessKey("cn-hangzhou", "LTAI4G7CSDhyEiXiVSk6zpnk", "pzvtq0LFSfF2EadlCxnZt4tLT4P4Em")
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Domain = "ecs.aliyuncs.com"
	request.Version = "2014-05-26"
	// 因为是RPC接口，因此需指定ApiName(Action)
	request.ApiName = "DescribeInstanceStatus"
	request.QueryParams["PageNumber"] = "1"
	request.QueryParams["PageSize"] = "40"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())
}
