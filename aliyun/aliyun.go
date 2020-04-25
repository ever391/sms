package aliyun

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func NewAliyunClient(signName, regionID, accessKeyID, accessSecret string) *aliyunClient {
	return &aliyunClient{
		SignName:     signName,
		RegionID:     regionID,
		AccessKeyID:  accessKeyID,
		AccessSecret: accessSecret,
	}
}

type aliyunClient struct {
	SignName     string
	RegionID     string
	AccessKeyID  string
	AccessSecret string
	ecsClient    *ecs.Client
}

func (c *aliyunClient) SendCode(params *params) (*ResponseMessage, error) {
	requestParams, err := c.BuildParams(params)
	if err != nil {
		return nil, err
	}
	if c.ecsClient == nil {
		c.ecsClient, err = ecs.NewClientWithAccessKey(c.RegionID, c.AccessKeyID, c.AccessSecret)
		if err != nil {
			return nil, err
		}
	}

	resp, err := c.ecsClient.ProcessCommonRequest(requestParams)
	if err != nil {
		// 异常处理
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, errors.New(fmt.Sprintf("status: %d, message: %s", resp.GetHttpStatus(), resp.GetHttpContentString()))
	}

	var respMsg = new(ResponseMessage)
	err = json.Unmarshal(resp.GetHttpContentBytes(), respMsg)
	if err != nil {
		return nil, err
	}
	respMsg.Mobile = params.Mobile
	return respMsg, nil
}

func (c *aliyunClient) BuildParams(params *params) (*requests.CommonRequest, error) {
	// 阿里短信服务SDK参数，采用硬编码
	var err error
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	// API 请求参数处理
	request.QueryParams["RegionId"] = c.RegionID
	request.QueryParams["SignName"] = c.SignName
	request.QueryParams["PhoneNumbers"] = params.Mobile
	request.QueryParams["TemplateCode"] = params.TemplateCode
	request.QueryParams["TemplateParam"], err = c.BuildTemplateParams(params.TemplateParams)
	return request, err
}

func (c *aliyunClient) BuildTemplateParams(templateParams map[string]string) (string, error) {
	// 短信内容参数处理
	resByte, err := json.Marshal(templateParams)
	if err != nil {
		return "", err
	}
	return string(resByte), nil
}

func (c *aliyunClient) CheckResponseMessage(respMsg *ResponseMessage) error {
	if respMsg.BizID == "" || respMsg.Message != "OK" || respMsg.Code != "OK" {
		err := errors.New(respMsg.Message)
		return err
	}
	return nil
}
