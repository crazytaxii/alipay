package alipay_test

import (
	"fmt"
	"testing"

	"github.com/crazytaxii/alipay"
)

func TestTradeRefund(t *testing.T) {
	alipayClient, err := alipay.NewClient(
		APP_ID,
		SELLER_ID,
		PEM_MY_PUB_KEY,
		PEM_MY_PVT_KEY,
		PEM_ALI_PUB_KEY,
	)
	if err != nil {
		fmt.Printf("alipay new client err: %s\n", err.Error())
	}

	tradeRefundRequestParam := alipay.TradeRefundRequestParam{
		OutTradeNo:   "20150320010101001",
		TradeNo:      "2014112611001004680073956707",
		RefundAmount: "0.01",
		RefundReason: "正常退款",
	}

	_, err = alipayClient.TradeRefund(tradeRefundRequestParam, alipay.SIGN_RSA2)
	if err != nil {
		fmt.Printf("alipay trade refund err: %s\n", err.Error())
	}
}
