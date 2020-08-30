package alipay

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	TransferRequestParam struct {
		OutBizNo      string `json:"out_base_no"`               // 商户转账唯一订单号
		PayeeType     string `json:"payee_type"`                // 收款方账户类型
		PayeeAccount  string `json:"payee_account"`             // 收款方账户类型
		Amount        string `json:"amount"`                    // 转账金额
		PayerShowName string `json:"payer_show_name,omitempty"` // 付款方姓名
		PayeeRealName string `json:"payee_real_name,omitempty"` // 收款方姓名
		Remark        string `json:"remark"`                    // 备注
	}

	TransferResponseParam struct {
		CommonResponseParam        // 公共响应参数
		OutBizNo            string `json:"out_biz_no"` // 商户转账唯一订单号
		OrderId             string `json:"order_id"`   // 支付宝转账单据号
		PayDate             string `json:"pay_date"`   // 支付时间
	}
)

func (c *AlipayClient) Transfer2Account(param *TransferRequestParam, signType string) (*TransferResponseParam, error) {
	bizContent, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	commonReqParam := &CommonRequestParam{
		AppID:      c.AppID,
		Method:     TRANSFER_2_ACCOUNT,
		Format:     "JSON",
		Charset:    "utf-8",
		SignType:   signType,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0",
		BizContent: string(bizContent),
	}
	respBody, err := c.doRequest(commonReqParam, signType)
	if err != nil {
		return nil, err
	}
	result := new(struct {
		TransferResParam TransferResponseParam `json:"alipay_fund_trans_toaccount_transfer_response"`
		Sign             string                `json:"sign"`
	})
	err = json.Unmarshal([]byte(respBody), result)
	if err != nil {
		return nil, err
	}
	if result.TransferResParam.Msg != "Success" {
		return nil, fmt.Errorf("alipay transfer to account failed, code: %s, sub_msg: %s, err_msg: %s",
			result.TransferResParam.Code, result.TransferResParam.SubCode, result.TransferResParam.SubMsg)
	}

	return &result.TransferResParam, nil
}

type (
	TransferQueryRequestParam struct {
		OutBizNo string `json:"out_biz_no,omitempty"` // 商户转账唯一订单号
		OrderId  string `json:"order_id,omitempty"`   // 支付宝转账单据号
	}

	TransferQueryResponseParam struct {
		CommonResponseParam        // 公共响应参数
		OrderId             string `json:"order_id"`         // 支付宝转账单据号
		Status              string `json:"status"`           // 转账单据状态
		PayDate             string `json:"pay_date"`         // 支付时间
		ArrivalTimeEnd      string `json:"arrival_time_end"` // 预计到账时间
		OrderFee            string `json:"order_fee"`        // 预计收费金额
		FailReason          string `json:"fail_reason"`      // 查询到的订单状态为FAIL失败或REFUND退票时，返回具体的原因
		OutBizNo            string `json:"out_biz_no"`       // 商户转账唯一订单号
		ErrCode             string `json:"err_code"`         // 错误代码
	}
)

func (c *AlipayClient) TransferQuery(param *TransferQueryRequestParam, signType string) (*TransferQueryResponseParam, error) {
	bizContent, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	commonReqParam := &CommonRequestParam{
		AppID:      c.AppID,
		Method:     TRANSFER_QUERY,
		Format:     "JSON",
		Charset:    "utf-8",
		SignType:   signType,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0",
		BizContent: string(bizContent),
	}
	respBody, err := c.doRequest(commonReqParam, signType)
	if err != nil {
		return nil, err
	}
	result := new(struct {
		TransferQueryResParam TransferQueryResponseParam `json:"alipay_fund_trans_order_query_response"`
		Sign                  string                     `json:"sign"`
	})
	err = json.Unmarshal([]byte(respBody), result)
	if err != nil {
		return nil, err
	}
	if result.TransferQueryResParam.Msg != "Success" {
		return nil, fmt.Errorf("alipay transfer to account failed, code: %s, sub_msg: %s, err_msg: %s",
			result.TransferQueryResParam.Code, result.TransferQueryResParam.SubCode, result.TransferQueryResParam.SubMsg)
	}

	return &result.TransferQueryResParam, nil
}
