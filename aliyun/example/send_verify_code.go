package main

import (
	"github.com/ever391/sms/aliyun"
)

func main() {
	var (
		SignName     = ""
		RegionID     = ""
		AccessKeyID  = ""
		AccessSecret = ""
		mobile       = "" // 发送短信目标手机号
		verifyCode   = "" // 业务生成的手机验证码
		product      = "" // 调用发送验证码的业务端的对外产品名称
		templateCode = "" // 阿里短信服务注册的短信模版
	)
	sms := aliyun.NewAliyunClient(SignName, RegionID, AccessKeyID, AccessSecret)
	params := aliyun.NewParams(mobile, verifyCode, product, templateCode)
	resp, err := sms.SendCode(params)
	if err != nil {
		// todo 自行处理错误
		return
	}
	err = sms.CheckResponseMessage(resp)
	if err != nil {
		// todo 自行处理错误
		return
	}
}
