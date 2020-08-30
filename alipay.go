package alipay

import (
	"crypto/rsa"
)

type AlipayClient struct {
	AppID        string
	SellerID     string
	PubKey       rsa.PublicKey
	PvtKey       rsa.PrivateKey
	AlipayPubKey rsa.PublicKey
}

var client *AlipayClient

func NewClient(appId, sellerId, myPubKey, myPvtKey, aliPubKey string) (*AlipayClient, error) {
	pubKey, err := parsePublicKey(myPubKey)
	if err != nil {
		return nil, err
	}
	pvtKey, err := parsePrivateKey(myPvtKey)
	if err != nil {
		return nil, err
	}
	alipayPubKey, err := parsePublicKey(aliPubKey)
	if err != nil {
		return nil, err
	}

	client = &AlipayClient{
		AppID:        appId,
		SellerID:     sellerId,
		PubKey:       *pubKey,
		PvtKey:       *pvtKey,
		AlipayPubKey: *alipayPubKey,
	}
	return client, nil
}

func GetClient() *AlipayClient {
	return client
}
