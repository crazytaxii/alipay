package alipay

import (
	"encoding/json"
	"time"

	"github.com/fatih/structs"
)

const (
	SIGN_RSA  = "RSA"
	SIGN_RSA2 = "RSA2"
)

type TradeAppPayRequestParam struct {
	TimeoutExpress string `json:"timeout_express,omitempty"` // 该笔订单允许的最晚付款时间，逾期将关闭交易。
	TotalAmount    string `json:"total_amount",omitempty`    // 订单总金额，单位为元，精确到小数点后两位
	SellerID       string `json:"seller_id,omitempty"`       // 收款支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	ProductCode    string `json:"product_code,omitempty"`    // 销售产品码，商家和支付宝签约的产品码
	Body           string `json:"body,omitempty"`            // 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
	Subject        string `json:"subject,omitempty"`         // 商品的标题/交易标题/订单标题/订单关键字等。
	OutTradeNo     string `json:"out_trade_no,omitempty"`    // 商户网站唯一订单号
	// TimeExpire     string `json:"time_expire,omitempty"`     // 绝对超时时间，格式为yyyy-MM-dd HH:mm。
	// GoodsType      string `json:"goods_type,omitempty"`      // 商品主类型 :0-虚拟类商品,1-实物类商品
	// PassbackParams string `json:"passback_params,omitempty"` // 公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。
	EnablePayChannels string `json:"enable_pay_channels,omitempty"` // 可用渠道，用户只能在指定渠道范围内支付
	// StoreId           string `json:"store_id,omitempty"`             // 商户门店编号
	// SpecifiedChannel  string `json:"specified_channel,omitempty"`    // 指定渠道，目前仅支持传入pcredit
	DisablePayChannel string `json:"disable_pay_channels,omitempty"` // 禁用渠道，用户不可用指定渠道支付
}

/**
 * app支付2.0
 */
func (c *AlipayClient) TradeAppPay(param TradeAppPayRequestParam,
	signType string, notifyURL string) (string, error) {
	bizContent, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	commonReqParam := &CommonRequestParam{
		AppID:      c.AppID,
		Method:     TRADE_APP_PAY,
		Format:     "JSON",
		Charset:    "utf-8",
		SignType:   signType,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0",
		NotifyURL:  notifyURL,
		BizContent: string(bizContent),
	}
	m := structs.Map(commonReqParam)
	sign, err := c.doSign(m, signType)
	if err != nil {
		return "", err
	}
	m["sign"] = sign

	return urlEncode(m), nil
} // TradeAppPay()
