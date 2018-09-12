package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
)

func doSignRSA(param map[string]interface{}, pvtKey *rsa.PrivateKey) (string, error) {
	pList := make([]string, 0, 0)
	for k, v := range param {
		if v != "" {
			pList = append(pList, fmt.Sprintf("%s=%s", k, strings.TrimSpace(fmt.Sprintf("%v", v))))
		}
	}
	sort.Strings(pList)
	src := strings.Join(pList, "&")
	sign, err := signPKCS1v15([]byte(src), pvtKey, crypto.SHA1)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
} // doSignRSA()

func doSignRSA2(param map[string]interface{}, pvtKey *rsa.PrivateKey) (string, error) {
	pList := make([]string, 0, 0)
	for k, v := range param {
		if v != "" {
			pList = append(pList, fmt.Sprintf("%s=%s", k, strings.TrimSpace(fmt.Sprintf("%v", v))))
		}
	}
	sort.Strings(pList)
	src := strings.Join(pList, "&")
	sign, err := signPKCS1v15([]byte(src), pvtKey, crypto.SHA256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
} // doSignRSA2()

func signPKCS1v15(src []byte, key *rsa.PrivateKey, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, key, hash, hashed)
} // signPKCS1v15()

func verifySignRSA(param map[string]interface{}, sign string, pubKey *rsa.PublicKey) bool {
	if sign == "" {
		return false
	}
	pList := make([]string, 0, 0)
	for k, v := range param {
		if k == "sign" {
			continue
		}
		if v != "" {
			pList = append(pList, fmt.Sprintf("%s=%s", k, strings.TrimSpace(fmt.Sprintf("%v", v))))
		}
	}
	sort.Strings(pList)
	src := strings.Join(pList, "&")
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}
	err = verifyPKCS1v15([]byte(src), sig, pubKey, crypto.SHA1)
	if err != nil {
		return false
	}
	return true
} // verifySignRSA()

func verifySignRSA2(param map[string]interface{}, sign string, pubKey *rsa.PublicKey) bool {
	if sign == "" {
		return false
	}
	pList := make([]string, 0, 0)
	for k, v := range param {
		if k == "sign" {
			continue
		}
		if v != "" {
			pList = append(pList, fmt.Sprintf("%s=%s", k, strings.TrimSpace(fmt.Sprintf("%v", v))))
		}
	}
	src := strings.Join(pList, "&")
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}
	err = verifyPKCS1v15([]byte(src), sig, pubKey, crypto.SHA256)
	if err != nil {
		return false
	}
	return true
} // verifySignRSA2()

func verifyPKCS1v15(src []byte, sign []byte, key *rsa.PublicKey, hash crypto.Hash) error {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	return rsa.VerifyPKCS1v15(key, hash, hashed, sign)
} // verifyPKCS1v15()
