package alipay_test

import (
	"fmt"

	"github.com/crazytaxii/alipay"
	"testing"
)

func TestTradeAppAay(t *testing.T) {
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

	tradeAppPayRequestParam := alipay.TradeAppPayRequestParam{
		TimeoutExpress: "30m",
		TotalAmount:    "0.01",
		SellerId:       SELLER_ID,
		ProductCode:    "QUICK_MSECURITY_PAY",
		Body:           "test body",
		Subject:        "test subject",
		OutTradeNo:     "70501111111S001111119",
	}
	payment, err := alipayClient.TradeAppPay(tradeAppPayRequestParam, alipay.SIGN_RSA2, "")
	if err != nil {
		fmt.Printf("trade app pay err: %s\n", err.Error())
	}
	fmt.Println("payment string:", payment)
}
