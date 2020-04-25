# 阿里短信服务 Aliyun SMS

### 快速开始

- 安装
```bash
go get github.com/ever391/sms/aliyun
```

- 发送短信
```
sms := aliyun.NewAliyunClient(SignName, RegionID, AccessKeyID, AccessSecret)
params := aliyun.NewParams(mobile, verifyCode, product, templateCode)
resp, err := sms.SendCode(params)
if err != nil {
    // todo 自行处理错误
    return
}
// 下边CheckResponseMessage可以省略，自己针对resp单独处理
err = sms.CheckResponseMessage(resp)
if err != nil {
    // todo 自行处理错误
    return
}
```
- 名词解释
    - SignName      签名，阿里服务
    - RegionID      地区，阿里服务 例：cn-hangzhou
    - AccessKeyID   授权KEY ID 阿里服务
    - AccessSecret  授权密钥 阿里服务
    - mobile        发送短信目标手机号
    - verifyCode    业务生成的手机验证码
    - product       调用发送验证码的业务端的对外产品名称
    - templateCode  阿里短信服务 短信模版 Code
    - templateParams阿里短信服务 短信模版 相应的参数健值对map[string]string{"code":"321123",...}
