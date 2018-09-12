package alipay

import ()

const BASE_URL = "https://openapi.alipay.com/gateway.do"

const (
	TRADE_APP_PAY = "alipay.trade.app.pay" // app支付接口2.0
	TRADE_REFUND  = "alipay.trade.refund"  // 统一收单交易退款
)

// 公共请求参数
type CommonRequestParam struct {
	AppId      string `structs:"app_id"`               // 支付宝分配给开发者的应用ID
	Method     string `structs:"method"`               // 接口名称
	Format     string `structs:"format,omitempty"`     // 仅支持JSON
	Charset    string `structs:"charset"`              // 请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType   string `structs:"sign_type"`            // 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	Timestamp  string `structs:"timestamp"`            // 发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version    string `structs:"version"`              // 调用的接口版本，固定为：1.0
	NotifyUrl  string `structs:"notify_url,omitempty"` // 支付宝服务器主动通知商户服务器里指定的页面http/https路径
	BizContent string `structs:"biz_content"`          // 请求参数的集合
}

// 公共响应参数
type CommonResponseParam struct {
	Code    string `json:"code" structs:"code"`                   // 网关返回码
	Msg     string `json:"msg" structs:"msg"`                     // 网关返回码描述
	SubCode string `json:"sub_code,omitempty" structs:"sub_code"` // 业务返回码
	SubMsg  string `json:"sub_msg,omitempty" structs:"sub_msg"`   // 业务返回码描述
}
