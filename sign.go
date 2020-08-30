package alipay

import (
	"bytes"
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"sort"
	"strings"
)

func (c *AlipayClient) doSign(param map[string]interface{}, signType string) (string, error) {
	pList := make([]string, 0, 0)
	for k, v := range param {
		if k == "sign" {
			// 剔除sign字段
			continue
		}
		if v != "" {
			pList = append(pList, fmt.Sprintf("%s=%s", k, strings.TrimSpace(fmt.Sprintf("%v", v))))
		}
	}
	sort.Strings(pList)
	src := strings.Join(pList, "&")
	buf := bytes.NewBufferString(src)
	var sign []byte
	var err error
	switch signType {
	case SIGN_RSA:
		sign, err = signPKCS1v15(buf.Bytes(), c.PvtKey, crypto.SHA1)
		if err != nil {
			return "", err
		}
	case SIGN_RSA2:
		sign, err = signPKCS1v15(buf.Bytes(), c.PvtKey, crypto.SHA256)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("Wrong sign type")
	}

	return base64.StdEncoding.EncodeToString(sign), nil
}

func (c *AlipayClient) doVerifySign(param map[string]interface{}, sign, signType string) bool {
	if sign == "" {
		return false
	}
	signB64, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}
	pList := make([]string, 0, 0)
	for k, v := range param {
		if k == "sign" || k == "sign_type" {
			// 剔除sign和sign_type两个参数
			continue
		}
		if v != "" {
			pList = append(pList, fmt.Sprintf("%s=%s", k, strings.TrimSpace(fmt.Sprintf("%v", v))))
		}
	}
	sort.Strings(pList)
	src := strings.Join(pList, "&")
	buf := bytes.NewBufferString(src)
	switch signType {
	case SIGN_RSA:
		err = verifyPKCS1v15(buf.Bytes(), signB64, c.AlipayPubKey, crypto.SHA1)
	case SIGN_RSA2:
		err = verifyPKCS1v15(buf.Bytes(), signB64, c.AlipayPubKey, crypto.SHA256)
	default:
		return false
	}
	if err != nil {
		return false
	}
	return true
}
