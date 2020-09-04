package main

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

/*
Project: aliyun images
Auth: gaopengcheng
Date: 2020-09-04 14:09
Version: v0.0.1
*/

// 定义镜像结构体
type Images struct {
	Data Data `json:"data"`
}
type Tags struct {
	Tag string `json:"tag"`
}
type Data struct {
	Tags []Tags `json:"tags"`
}

// 阿里云认证
func ailynAuth() (res string) {
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", "LTAI4GKuRVd8TtNift14gvek", "1fvsP63srlqMcZWuJhcLB72IMHgXNw")
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Scheme = "https" // https | http
	request.Domain = "cr.cn-hangzhou.aliyuncs.com"
	request.Version = "2016-06-07"
	request.PathPattern = "/repos/guming_teste/book-starter-client/tags"
	request.Headers["Content-Type"] = "application/json"

	body := `{}`
	request.Content = []byte(body)

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	tagInfo := response.GetHttpContentString()
	return tagInfo
}

// 处理镜像数据
func imageDate(b []byte) {
	var images Images
	err := json.Unmarshal(b, &images)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(images.Data.Tags[3].Tag)
	for i, _ := range images.Data.Tags {
		tagVersion := images.Data.Tags[i].Tag
		fmt.Println(tagVersion)
		if i == 10 {
			break
		}
	}
}

func main() {
	tagInfo1 := ailynAuth()
	imageDate([]byte(tagInfo1))
}
