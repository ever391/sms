package aliyun

import (
	"errors"
	"regexp"
)

func NewParams(mobile, verifyCode, product, templateCode string) *params {
	return &params{
		Mobile:         mobile,
		VerifyCode:     verifyCode,
		Product:        product,
		TemplateCode:   templateCode,
		TemplateParams: map[string]string{"code": verifyCode},
	}
}

type params struct {
	Mobile         string
	VerifyCode     string
	Product        string
	TemplateCode   string
	TemplateParams map[string]string
}

func (p *params) CheckParamsAll() error {
	if p.CheckMobile() == false {
		return errors.New("mobile is wrong")
	}
	if p.CheckVerifyCode() == false {
		return errors.New("mobile is wrong")
	}
	if p.CheckTemplateCode() == false {
		return errors.New("mobile is wrong")
	}
	if p.CheckProduct() == false {
		return errors.New("mobile is wrong")
	}

	return nil
}

func (p *params) CheckMobile() bool {
	re, _ := regexp.Compile(`^1[3-9]\d{9}$`)
	return re.MatchString(p.Mobile)
}

func (p *params) CheckVerifyCode() bool {
	return len(p.VerifyCode) != 0
}

func (p *params) CheckTemplateCode() bool {
	return p.TemplateCode != ""
}

func (p *params) CheckProduct() bool {
	return p.Product != ""
}
