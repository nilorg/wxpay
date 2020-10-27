package wxpay

import (
	"encoding/xml"

	"github.com/nilorg/sdk/random"
)

// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2

// PromotionTransfersRequest 企业向微信用户个人付款请求
type PromotionTransfersRequest struct {
	XMLName        xml.Name `xml:"xml"`
	MchAppID       string   `xml:"mch_appid"`
	MchID          string   `xml:"mchid"`
	DeviceInfo     string   `xml:"device_info"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	OpenID         string   `xml:"openid"`
	CheckName      string   `xml:"check_name"`
	ReUserName     string   `xml:"re_user_name"`
	Amount         int      `xml:"amount"`
	Desc           string   `xml:"desc"`
	SpbillCreateIP string   `xml:"spbill_create_ip"`
}

// NewPromotionTransfersRequest 创建企业向微信用户个人付款请求
func NewPromotionTransfersRequest() *PromotionTransfersRequest {
	return &PromotionTransfersRequest{
		NonceStr: random.AZaz09(32),
	}
}

// SignMD5 md5
func (req *PromotionTransfersRequest) SignMD5(apiKey string) error {
	params, err := SignStructToParameter(*req)
	if err != nil {
		return err
	}
	params["key"] = apiKey
	value := SignMD5(params)
	req.Sign = value
	return nil
}

// PromotionTransfersResponse 企业向微信用户个人付款响应
type PromotionTransfersResponse struct {
	ResponseStatus
	MchAppID       string `xml:"mch_appid"`
	MchID          string `xml:"mchid"`
	DeviceInfo     string `xml:"device_info"`
	NonceStr       string `xml:"nonce_str"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	PaymentNo      string `xml:"payment_no"`
	PaymentTime    string `xml:"payment_time"`
}
