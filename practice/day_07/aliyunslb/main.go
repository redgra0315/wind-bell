package main

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

type AutoGenerated struct {
	BackendServers BackendServers `json:"BackendServers"`
}
type BackendServer struct {
	Type     string `json:"Type"`
	ServerID string `json:"ServerId"`
	Port     int    `json:"Port"`
	Weight   int    `json:"Weight"`
}
type BackendServers struct {
	BackendServer []BackendServer `json:"BackendServer"`
}

// aliyun slb auth
func ailsunAuth() {
	client, err := slb.NewClientWithAccessKey("cn-hangzhou", "", "")
	request := slb.CreateSetVServerGroupAttributeRequest()
	request.Scheme = "https"

	request.VServerGroupId = "rsp-bp1njbb03c6uq"
	request.VServerGroupName = "test-xueyuan"

	//request.BackendServers = "[{\"Type\":\"ecs\",\"ServerId\":\"i-bp1gqo4j4rk88zkr7bk0\",\"Port\":80,\"Weight\":100}]"

	response, err := client.SetVServerGroupAttribute(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	slnInfo := response.BaseResponse.GetHttpContentString()

	var serverGroup AutoGenerated
	json.Unmarshal([]byte(slnInfo), &serverGroup)
	slnInfo1 := serverGroup.BackendServers.BackendServer[0]
	//fmt.Println(slnInfo.Type)
	//fmt.Println(slnInfo.ServerID)
	//fmt.Println(slnInfo.Port)
	slnInfo1.Weight = 70
	aa := NewBackendServer(slnInfo1.Type, slnInfo1.ServerID, slnInfo1.Port, slnInfo1.Weight)
	fmt.Println(aa)

	request.BackendServers = fmt.Sprintf("[{\"Type\":\"ecs\",\"ServerId\":\"i-bp1gqo4j4rk88zkr7bk0\",\"Port\":\"80\",\"Weight\": %d}]", slnInfo1.Weight)

	//request.BackendServers = aa
	response, err = client.SetVServerGroupAttribute(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// 构造 BackendServer 结构体
func NewBackendServer(Type, ServerID string, Port, Weight int) *BackendServer {
	return &BackendServer{
		Type:     Type,
		ServerID: ServerID,
		Port:     Port,
		Weight:   Weight,
	}
}

// update aliyun slb weigh

func main() {
	ailsunAuth()

}
