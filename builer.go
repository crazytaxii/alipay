package alipay

import (
	"errors"
	"fmt"

	"github.com/fatih/structs"
)

func (alipayClient *AlipayClient) doRequest(param interface{}, signType string) (string, error) {
	m := structs.Map(param)
	var sign string
	var err error
	if signType == SIGN_RSA {
		sign, err = doSignRSA(m, alipayClient.MyPvtKey)
	} else if signType == SIGN_RSA2 {
		sign, err = doSignRSA2(m, alipayClient.MyPvtKey)
	} else {
		return "", errors.New("Wrong sign type")
	}
	if err != nil {
		return "", err
	}
	if sign == "" {
		return "", errors.New("Do sign failed")
	}
	m["sign"] = sign

	resp, body, errs := alipayClient.Agent.Get(fmt.Sprintf("%s?%s", BASE_URL, urlEncode(m))).
		Set("Content-Type", "application/json").End()
	if errs != nil {
		return "", errs[0]
	}
	fmt.Println("debug: resp status:", resp.Status)
	fmt.Println("debug: resp body:", body)

	return body, nil
} // doRequest()
