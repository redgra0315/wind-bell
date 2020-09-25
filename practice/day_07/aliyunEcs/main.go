package main

/*
ECS 获取IP地址
1  根据slb的id来查找到IP地址
*/

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
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

// 切片转json
func jsonString() string {
	type EcsId []string
	aa := EcsId{"i-bp1gqo4j4rk88zkr7bk0", "i-bp17uhlzld7877uh3bs4"}
	//fmt.Println(aa)
	aaa, _ := json.Marshal(aa)
	a := string(aaa)
	return a
}

func aliyuAuth(a string) (res string) {
	//Config := sdk.NewConfig().
	//	WithEnableAsync(true).
	//	WithGoRoutinePoolSize(5).
	//	WithMaxTaskQueueSize(1000)
	//credential := &credentials.BaseCredential{
	//	AccessKeyId: "",
	//	AccessKeySecret: "",
	//}

	//client, err := sdk.NewClientWithOptions("cn-hangzhou",Config,credential)
	cca, _ := sdk.NewClientWithBearerToken()

	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", "", "")

	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"

	//ecsID := ["i-bp1gqo4j4rk88zkr7bk0"]

	//request.InstanceIds = "[\"i-bp1gqo4j4rk88zkr7bk0\"]"
	request.InstanceIds = a

	response, err := client.DescribeInstances(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	// 直接获取string
	res = response.BaseResponse.GetHttpContentString()
	//fmt.Println(res)
	return res

	//fmt.Println(res)

	//ipAdder, _ := gojsonq.New().JSONString(res).FindR("Instances.Instance")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//name,_ := ipAdder.String()
	//fmt.Printf("%#v\n", name)

	//jq := gojsonq.New().File(res).From("Instances").Select("RequestId", "TotalCount").WhereNotNil("RequestId")
	//fmt.Printf("%#v\n", jq.Get())

	// 获取的值为byte，byte转json
	//res1 := response.BaseResponse.GetHttpContentBytes()
	//
	//return res1

	//js, err := simplejson.NewJson([]byte(res1))
	//if err != nil {
	//	panic(err.Error())
	//}
	//personArr, err := js.Get("Instances").Get("Instance").Get("VpcAttributes").Get("PrivateIpAddress").Get("IpAddress")
	//fmt.Println(personArr)
}

func main() {
	//bb := Slice()
	//fmt.Println(bb)
	//aa := aliyuAuth("i-bp1gqo4j4rk88zkr7bk0")
	easedInfo := jsonString()
	aa := aliyuAuth(easedInfo)
	method2([]byte(aa))
	//fmt.Println(aa)
	//b, err := ioutil.ReadFile(aa)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//method1([]byte(b))
	////method2(b)
	//
	////aa := aliyuAuth()
	////fmt.Println(aa)
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
	for i := 0; i < 2; i++ {
		fmt.Println(Instance[i].VpcAttributes.PrivateIpAddress.IpAddress[0])
	}
}
