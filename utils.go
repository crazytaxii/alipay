package alipay

import (
	"crypto"
	"crypto/rand"
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
func parsePublicKey(pubPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPem))
	if block == nil {
		return nil, errors.New("Failed to parse PEM block containing the public key")
	}
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
}

/**
 * 解析私钥证书
 */
func parsePrivateKey(pvtPem string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pvtPem))
	if block == nil {
		return nil, errors.New("Failed to parse PEM block containing the private key")
	}
	pvtKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse DER encoded private key, err: %s",
			err.Error())
	}
	return pvtKey, nil
}

func urlEncode(param map[string]interface{}) string {
	urlValues := &url.Values{}
	for k, v := range param {
		urlValues.Add(k, fmt.Sprintf("%v", v))
	}
	return urlValues.Encode()
}

func signPKCS1v15(src []byte, key rsa.PrivateKey, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, &key, hash, hashed)
}

func verifyPKCS1v15(src, sign []byte, key rsa.PublicKey, hash crypto.Hash) error {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	return rsa.VerifyPKCS1v15(&key, hash, hashed, sign)
}
