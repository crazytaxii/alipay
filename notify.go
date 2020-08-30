package alipay

import (
	"errors"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gorilla/schema"
)

type TradeNotificationParam struct {
	NotifyTime        string `structs:"notify_time" schema:"notify_time,required"` // 通知时间
	NotifyType        string `structs:"notify_type" schema:"notify_type,required"` // 通知类型
	NotifyID          string `structs:"notify_id" schema:"notify_id,required"`     // 通知校验ID
	AppID             string `structs:"app_id" schema:"app_id,required"`           // 支付宝分配给开发者的应用ID
	AuthAppID         string `structs:"auth_app_id" schema:"auth_app_id,required"`
	Charset           string `structs:"charset" schema:"charset,required"`                          // 编码格式
	Version           string `structs:"version" schema:"version,required"`                          // 接口版本
	SignType          string `structs:"sign_type" schema:"sign_type,required"`                      // 签名类型
	Sign              string `structs:"sign" schema:"sign,required"`                                // 签名
	TradeNo           string `structs:"trade_no" schema:"trade_no,required"`                        // 支付宝交易号
	OutTradeNo        string `structs:"out_trade_no" schema:"out_trade_no,required"`                // 商户订单号
	OutBizNo          string `structs:"out_biz_no,omitempty" schema:"out_biz_no"`                   // 商户业务号
	BuyerID           string `structs:"buyer_id,omitempty" schema:"buyer_id"`                       // 买家支付宝用户号
	BuyerLogonID      string `structs:"buyer_logon_id,omitempty" schema:"buyer_logon_id"`           // 买家支付宝账号
	SellerID          string `structs:"seller_id,omitempty" schema:"seller_id"`                     // 卖家支付宝用户号
	SellerEmail       string `structs:"seller_email,omitempty" schema:"seller_email"`               // 卖家支付宝账号
	TradeStatus       string `structs:"trade_status,omitempty" schema:"trade_status"`               // 交易状态
	TotalAmount       string `structs:"total_amount,omitempty" schema:"total_amount"`               // 订单金额
	ReceiptAmount     string `structs:"receipt_amount,omitempty" schema:"receipt_amount"`           // 实收金额
	InvoiceAmount     string `structs:"invoice_amount,omitempty" schema:"invoice_amount"`           // 开票金额
	BuyerPayAmount    string `structs:"buyer_pay_amount,omitempty" schema:"buyer_pay_amount"`       // 付款金额
	PointAmount       string `structs:"point_amount,omitempty" schema:"point_amount"`               // 集分宝金额
	RefundFee         string `structs:"refund_fee,omitempty" schema:"refund_fee"`                   // 总退款金额
	Subject           string `structs:"subject,omitempty" schema:"subject"`                         // 订单标题
	Body              string `structs:"body,omitempty" schema:"body"`                               // 商品描述
	GmtCreate         string `structs:"gmt_create,omitempty" schema:"gmt_create"`                   // 交易创建时间
	GmtPayment        string `structs:"gmt_payment,omitempty" schema:"gmt_payment"`                 // 交易付款时间
	GmtRefund         string `structs:"gmt_refund,omitempty" schema:"gmt_refund"`                   // 交易退款时间
	GmtClose          string `structs:"gmt_close,omitempty" schema:"gmt_close"`                     // 交易结束时间
	FundBillList      string `structs:"fund_bill_list,omitempty" schema:"fund_bill_list"`           // 支付金额信息
	PassbackParams    string `structs:"passback_params,omitempty" schema:"passback_params"`         // 回传参数
	VoucherDetailList string `structs:"voucher_detail_list,omitempty" schema:"voucher_detail_list"` // 优惠券信息
}

/**
 * 解析支付宝通知请求并验签
 */
func (c *AlipayClient) GetTradeNotification(req *http.Request) (*TradeNotificationParam, error) {
	if req == nil {
		return nil, errors.New("Nil request")
	}
	err := req.ParseForm()
	if err != nil {
		return nil, err
	}

	param := new(TradeNotificationParam)
	err = schema.NewDecoder().Decode(param, req.PostForm)
	if err != nil {
		return nil, err
	}

	if ok := c.doVerifySign(structs.Map(param), param.Sign, param.SignType); !ok {
		return nil, errors.New("Verify sign failed")
	}

	return param, nil
}
