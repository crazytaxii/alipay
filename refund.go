package alipay

import (
	"encoding/json"
	"fmt"
	"time"
)

type TradeRefundRequestParam struct {
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

type TradeRefundResponseParam struct {
	CommonResponseParam
	// Code    string `json:"code" structs:"code"`                   // 网关返回码
	// Msg     string `json:"msg" structs:"msg"`                     // 网关返回码描述
	// SubCode string `json:"sub_code,omitempty" structs:"sub_code"` // 业务返回码
	// SubMsg  string `json:"sub_msg,omitempty" structs:"sub_msg"`   // 业务返回码描述

	TradeNo        string `json:"trade_no" structs:"trade_no,omitempty"`             // 支付宝交易号
	OutTradeNo     string `json:"out_trade_no" structs:"out_trade_no,omitempty"`     // 商户订单号
	BuyerLogonId   string `json:"buyer_logon_id" structs:"buyer_logon_id,omitempty"` // 用户的登录ID
	FundChange     string `json:"fund_change" structs:"fund_change,omitempty"`       // 本次退款是否发生了资金变
	RefundFee      string `json:"refund_fee" structs:"refund_fee,omitempty"`         // 退款总金额
	SendBackFee    string `json:"send_back_fee" structs:"send_back_fee",omitempty`
	RefundCurrency string `json:"refund_currency" structs:"refund_currency,omitempty"` // 退款币种信息
	GmtRefundPay   string `json:"gmt_refund_pay" structs:"gmt_refund_pay,omitempty"`   // 退款使用的资金渠道
	StoreName      string `json:"store_name,omitempty" structs:"store_name,omitempty"` // 交易在支付时候的门店名称
	BuyerUserId    string `json:"buyer_user_id" structs:"buyer_user_id,omitempty"`     // 买家在支付宝的用户ID
	// PresentRefundBuyerAmount     string `json:"present_refund_buyer_amount,omitempty" structs:""`     // 本次退款金额中买家退款金额
	// PresentRefundDiscountAmount  string `json:"present_refund_discount_amount,omitempty" structs:""`  // 本次退款金额中平台优惠退款金额
	// PresentRefundMdiscountAmount string `json:"present_refund_mdiscount_amount,omitempty" structs:""` // 本次退款金额中商家优惠退款金额
}

/**
 * 退款
 */
func (alipayClient *AlipayClient) TradeRefund(tradeRefundRequestParam TradeRefundRequestParam,
	signType string) (*TradeRefundResponseParam, error) {
	bizContent, err := json.Marshal(tradeRefundRequestParam)
	if err != nil {
		return nil, err
	}

	requestTime := time.Now().Format("2006-01-02 15:04:05")
	commonRequestParam := &CommonRequestParam{
		AppId:      alipayClient.AppId,
		Method:     TRADE_REFUND,
		Format:     "JSON",
		Charset:    "utf-8",
		SignType:   signType,
		Timestamp:  requestTime,
		Version:    "1.0",
		BizContent: string(bizContent),
	}
	respBody, err := alipayClient.doRequest(commonRequestParam, signType)
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
