package wxpay

import (
	"encoding/xml"
	"time"

	"github.com/nilorg/sdk/random"
)

// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1

// UnifiedOrderRequest 统一支付
type UnifiedOrderRequest struct {
	XMLName        xml.Name `xml:"xml"`
	AppID          string   `xml:"appid"`
	MchID          string   `xml:"mch_id"`
	DeviceInfo     string   `xml:"device_info"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	SignType       string   `xml:"sign_type"`
	Body           string   `xml:"body"`
	Detail         CDATA    `xml:"detail"`
	Attach         string   `xml:"attach"`
	OutTradeNo     string   `xml:"out_trade_no"`
	FeeType        string   `xml:"fee_type"`
	TotalFee       uint64   `xml:"total_fee"`
	SpbillCreateIP string   `xml:"spbill_create_ip"`
	TimeStart      string   `xml:"time_start"`
	TimeExpire     string   `xml:"time_expire"`
	GoodsTag       string   `xml:"goods_tag"`
	NotifyURL      string   `xml:"notify_url"`
	TradeType      string   `xml:"trade_type"`
	ProductID      string   `xml:"product_id"`
	LimitPay       string   `xml:"limit_pay"`
	OpenID         string   `xml:"openid"`
	Receipt        string   `xml:"receipt"`
	SceneInfo      string   `xml:"scene_info"`
}

// NewUnifiedOrderRequest 创建统一支付
func NewUnifiedOrderRequest() *UnifiedOrderRequest {
	timeStart := time.Now()
	return &UnifiedOrderRequest{
		TimeStart:  timeStart.Format("20060102150405"),
		TimeExpire: timeStart.Add(time.Minute * 30).Format("20060102150405"),
		NonceStr:   random.AZaz09(32),
	}
}

// SignMD5 md5
func (req *UnifiedOrderRequest) SignMD5(apiKey string) error {
	req.SignType = MD5
	params, err := SignStructToParameter(*req)
	if err != nil {
		return err
	}
	params["key"] = apiKey
	value := SignMD5(params)
	req.Sign = value
	return nil
}

// SignHMACSHA256 HMACS-HA256
func (req *UnifiedOrderRequest) SignHMACSHA256(key []byte, apiKey string) error {
	req.SignType = HMACSHA256
	params, err := SignStructToParameter(*req)
	if err != nil {
		return err
	}
	params["key"] = apiKey
	value := SignHMACSHA256(params, key)
	req.Sign = value
	return nil
}

// UnifiedOrderResponse ...
type UnifiedOrderResponse struct {
	ResponseStatus
	TradeType string `xml:"trade_type"`
	PrepayID  string `xml:"prepay_id"`
	CodeURL   string `xml:"code_url"`
}
