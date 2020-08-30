package alipay

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	TradeRefundRequestParam struct {
		OutTradeNo     string `json:"out_trade_no,omitempty"`    // 订单支付时传入的商户订单号,不能和 trade_no同时为空。
		TradeNo        string `json:"trade_no,omitempty"`        // 支付宝交易号，和商户订单号不能同时为空
		RefundAmount   string `json:"refund_amount"`             // 需要退款的金额
		RefundCurrency string `json:"refund_currency,omitempty"` // 订单退款币种信息
		RefundReason   string `json:"refund_reason,omitempty"`   // 退款的原因说明
		OutRequestNo   string `json:"out_request_no,omitempty"`  // 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
		OperatorId     string `json:"operator_id,omitempty"`     // 商户的操作员编号
		StoreId        string `json:"store_id,omitempty"`        // 商户的门店编号
		TerminalId     string `json:"terminal_id,omitempty"`     // 商户的终端编号
	}

	TradeRefundResponseParam struct {
		CommonResponseParam        // 公共响应参数
		TradeNo             string `json:"trade_no" structs:"trade_no,omitempty"`               // 支付宝交易号
		OutTradeNo          string `json:"out_trade_no" structs:"out_trade_no,omitempty"`       // 商户订单号
		BuyerLogonID        string `json:"buyer_logon_id" structs:"buyer_logon_id,omitempty"`   // 用户的登录ID
		FundChange          string `json:"fund_change" structs:"fund_change,omitempty"`         // 本次退款是否发生了资金变
		RefundFee           string `json:"refund_fee" structs:"refund_fee,omitempty"`           // 退款总金额
		RefundCurrency      string `json:"refund_currency" structs:"refund_currency,omitempty"` // 退款币种信息
		GmtRefundPay        string `json:"gmt_refund_pay" structs:"gmt_refund_pay,omitempty"`   // 退款使用的资金渠道
		StoreName           string `json:"store_name,omitempty" structs:"store_name,omitempty"` // 交易在支付时候的门店名称
		BuyerUserID         string `json:"buyer_user_id" structs:"buyer_user_id,omitempty"`     // 买家在支付宝的用户ID
	}
)

/**
 * 退款
 */
func (c *AlipayClient) TradeRefund(tradeRefundRequestParam TradeRefundRequestParam,
	signType string) (*TradeRefundResponseParam, error) {
	bizContent, err := json.Marshal(tradeRefundRequestParam)
	if err != nil {
		return nil, err
	}

	commonRequestParam := &CommonRequestParam{
		AppID:      c.AppID,
		Method:     TRADE_REFUND,
		Format:     "JSON",
		Charset:    "utf-8",
		SignType:   signType,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0",
		BizContent: string(bizContent),
	}
	respBody, err := c.doRequest(commonRequestParam, signType)
	if err != nil {
		return nil, err
	}
	result := new(struct {
		TradeRefundResponseParam TradeRefundResponseParam `json:"alipay_trade_refund_response"`
		Sign                     string                   `json:"sign"`
	})
	err = json.Unmarshal([]byte(respBody), result)
	if err != nil {
		return nil, err
	}
	if result.TradeRefundResponseParam.Msg != "Success" {
		return nil, fmt.Errorf("Alipay trade refund failed, code: %s, sub_code: %s, err_msg: %s",
			result.TradeRefundResponseParam.Code, result.TradeRefundResponseParam.SubCode,
			result.TradeRefundResponseParam.SubMsg)
	}

	return &result.TradeRefundResponseParam, nil
} // TradeRefund()

type (
	TradeRefundQueryRequestParam struct {
		TradeNo      string `json:"trade_no",omitempty`     // 支付宝交易号，和商户订单号不能同时为空
		OutTradeNo   string `json:"out_trade_no",omitempty` // 订单支付时传入的商户订单号,和支付宝交易号不能同时为空
		OutRequestNo string `json:"out_request_no"`         // 请求退款接口时，传入的退款请求号
	}

	TradeRefundQueryResponseParam struct {
		CommonResponseParam // 公共响应参数

		TradeNo                      string `json:"trade_no" structs:"trade_no,omitempty"`
		OutTradeNo                   string `json:"out_trade_no" structs:"out_trade_no,omitempty"`
		OutRequestNo                 string `json:"out_request_no" structs:"out_request_no,omitempty"`
		RefundReason                 string `json:"refund_reason" structs:"refund_reason,omitempty"`
		TotalAmount                  string `json:"total_amount" structs:"total_amount,omitempty"`
		RefundAmount                 string `json:"refund_amount" structs:"refund_amount,omitempty"`
		PresentRefundDiscountAmount  string `json:"present_refund_discount_amount" structs:"present_refund_discount_amount,omitempty"`
		PresentRefundMdiscountAmount string `json:"present_refund_mdiscount_amount" structs:"present_refund_mdiscount_amount,omitempty"`
	}
)

/**
 * 退款查询
 */
func (c *AlipayClient) TradeRefundQuery(tradeRefundQueryRequestParam *TradeRefundQueryRequestParam,
	signType string) (*TradeRefundQueryResponseParam, error) {
	bizContent, err := json.Marshal(tradeRefundQueryRequestParam)
	if err != nil {
		return nil, err
	}

	commonRequestParam := &CommonRequestParam{
		AppID:      c.AppID,
		Method:     TRADE_REFUND_QUERY,
		Format:     "JSON",
		Charset:    "utf-8",
		SignType:   signType,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0",
		BizContent: string(bizContent),
	}
	respBody, err := c.doRequest(commonRequestParam, signType)
	if err != nil {
		return nil, err
	}
	result := new(struct {
		TradeRefundQueryResponseParam `json:"alipay_trade_fastpay_refund_query_response"`
		Sign                          string `json:"sign"`
	})
	err = json.Unmarshal([]byte(respBody), result)
	if err != nil {
		return nil, err
	}
	if result.TradeRefundQueryResponseParam.Msg != "Success" {
		return nil, fmt.Errorf("Alipay trade refund query failed, code: %s, sub_code: %s, err_msg: %s",
			result.TradeRefundQueryResponseParam.Code, result.TradeRefundQueryResponseParam.SubCode,
			result.TradeRefundQueryResponseParam.SubMsg)
	}

	return &result.TradeRefundQueryResponseParam, nil
} // TradeRefundQuery()
