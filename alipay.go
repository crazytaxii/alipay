package alipay

import (
	"crypto/rsa"

	"github.com/parnurzeal/gorequest"
)

type AlipayClient struct {
	AppId        string
	SellerId     string
	MyPubKey     *rsa.PublicKey
	MyPvtKey     *rsa.PrivateKey
	AlipayPubKey *rsa.PublicKey
	Agent        *gorequest.SuperAgent
}

var alipayClient = new(AlipayClient)

func NewClient(appId string, sellerId string, myPubKey string, myPvtKey string, aliPubKey string) (*AlipayClient, error) {
	alipayClient.AppId = appId
	alipayClient.SellerId = sellerId
	var err error
	alipayClient.MyPubKey, err = ParsePublicKey(myPubKey)
	if err != nil {
		return nil, err
	}
	alipayClient.MyPvtKey, err = ParsePrivateKey(myPvtKey)
	if err != nil {
		return nil, err
	}
	alipayClient.AlipayPubKey, err = ParsePublicKey(aliPubKey)
	if err != nil {
		return nil, err
	}
	alipayClient.Agent = gorequest.New()
	return alipayClient, nil
} // NewClient()

func GetClient() *AlipayClient {
	return alipayClient
} // GetClient()
