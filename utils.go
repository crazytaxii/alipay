package alipay

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
)

/**
 * 解析公钥证书
 */
func ParsePublicKey(pubPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPem))
	if block == nil {
		return nil, errors.New("Failed to parse PEM block containing the public key")
	}
	// fmt.Println("block.Type:", block.Type)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse DER encoded public key, err: %s",
			err.Error())
	}
	pubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Unable to parse RSA public key")
	}
	return pubKey, nil
} // ParsePublicKey()

/**
 * 解析私钥证书
 */
func ParsePrivateKey(pvtPem string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pvtPem))
	if block == nil {
		return nil, errors.New("Failed to parse PEM block containing the private key")
	}
	// fmt.Println("block.Type:", block.Type)
	pvtKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse DER encoded private key, err: %s",
			err.Error())
	}
	return pvtKey, nil
} // ParsePrivateKey()

func urlEncode(param map[string]interface{}) string {
	urlValues := &url.Values{}
	for k, v := range param {
		urlValues.Add(k, fmt.Sprintf("%v", v))
	}
	return urlValues.Encode()
} // urlEncode()
