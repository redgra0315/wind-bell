package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

func demo(aa string) {
	Config := sdk.NewConfig().
		WithEnableAsync(true).
		WithGoRoutinePoolSize(5).
		WithMaxTaskQueueSize(1000)
	credential := &credentials.BaseCredential{
		AccessKeyId:     "",
		AccessKeySecret: "",
	}
	client, err := sdk.NewClientWithOptions("cn-hangzhou", Config, credential)
	if err != nil {
		panic(err)
	}
	//request := slb.CreateSetBackendServersRequest()
	//request.Scheme = "https"
	//request.LoadBalancerId = "rsp-bp1njbb03c6uq"
	//request.BackendServers = "test-xueyuan"
	//
	//
	//
	//
	//
	//response, err := slb.CreateSetBackendServersRequest(request)
	//if err != nil {
	//	fmt.Print(err.Error())
	//}
	//fmt.Printf("response is %#v\n", response)

	request := slb.CreateSetVServerGroupAttributeRequest()
	request.Scheme = "https"

	request.VServerGroupId = "rsp-bp1njbb03c6uq"
	request.VServerGroupName = "test-xueyuan"

	//request.BackendServers = "[{\"Type\":\"ecs\",\"ServerId\":\"i-bp1gqo4j4rk88zkr7bk0\",\"Port\":80,\"Weight\":100}]"

	request.BackendServers = aa
	response, err := client.SetVServerGroupAttribute(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	slnInfo = response.BaseResponse.GetHttpContentString()
	return slnInfo

}

func main() {
	slb.CreateSetBackendServersRequest()
	demo()

}
