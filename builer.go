package alipay

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/parnurzeal/gorequest"
)

func (c *AlipayClient) doRequest(param interface{}, signType string) (string, error) {
	m := structs.Map(param)
	sign, err := c.doSign(m, signType)
	if err != nil {
		return "", err
	}
	m["sign"] = sign

	_, body, errs := gorequest.New().Get(fmt.Sprintf("%s?%s", BASE_URL, urlEncode(m))).
		Set("Content-Type", "application/json").End()
	if errs != nil {
		return "", errs[0]
	}

	return body, nil
} // doRequest()
