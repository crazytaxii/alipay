package alipay

import (
	"encoding/json"
	"fmt"
	"time"
)

type AntiFraudRequestParam struct {
	ProductCode   string `json:"product_code,required"`   // 产品码，标记商户接入的具体产品
	TransactionId string `json:"transaction_id,required"` // 商户请求的唯一标志
	CertNo        string `json:"cert_no,required"`        // 证件号
	CertType      string `json:"cert_type,required"`      // 证件类型
	Name          string `json:"name,required"`           // 姓名
	Mobile        string `json:"mobile,omitempty"`        // 手机号码
	Email         string `json:"email,omitempty"`         // 电子邮箱
	BankCard      string `json:"bank_card,omitempty"`     // 仅支持大陆银行卡验证，包括信用卡、借记卡等实体卡。
	Address       string `json:"address,omitempty"`       // 地址信息
	// Ip            string `json:"ip,omitempty"`            // ip地址
	// Mac           string `json:"mac,omitempty"`           // 物理地址
	// Wifimac       string `json:"wifimac,omitempty"`       // wifi的物理地址
	// Imei          string `json:"imei,omitempty"`          // 国际移动设备标志
}

type AntiFraudResponseParam struct {
	CommonResponseParam

	BizNo          string `json:"biz_no"`
	VerifyCode     string `json:"verify_code"`
	SolutionId     string `json:"solution_id"`
	DecisionResult string `json:"decision_result"`
}

/**
 * 欺诈信息验证
 */
func (alipayClient *AlipayClient) AntiFraudVerify(antiFraudRequestParam *AntiFraudRequestParam,
	signType string) error {
	bizContent, err := json.Marshal(antiFraudRequestParam)
	if err != nil {
		return err
	}

	commonRequestParam := &CommonRequestParam{
		AppId:      alipayClient.AppId,
		Method:     TRADE_REFUND,
		Format:     "JSON",
		Charset:    "utf-8",
		SignType:   signType,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0",
		BizContent: string(bizContent),
	}
	respBody, err := alipayClient.doRequest(commonRequestParam, signType)
	if err != nil {
		return err
	}
	result := new(struct {
		AntiFraudResponseParam AntiFraudResponseParam `json:"zhima_credit_antifraud_verify_response"`
		Sign                   string                 `json:"sign"`
	})
	err = json.Unmarshal([]byte(respBody), result)
	if err != nil {
		return err
	}
	if result.AntiFraudResponseParam.Msg != "Success" {
		return fmt.Errorf("Zhima anti fraud verified failed, code: %s, sub_code: %s, err_msg: %s",
			result.AntiFraudResponseParam.Code, result.AntiFraudResponseParam.SubCode,
			result.AntiFraudResponseParam.SubMsg)
	}
	if result.AntiFraudResponseParam.DecisionResult != "PASS" {
		return fmt.Errorf("Zhima anti fraud verified failed")
	}
	return nil
} // AntiFraudVerify()
