package wxpay

import (
	"encoding/xml"

	"github.com/nilorg/sdk/random"
)

// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_5&index=4

// SendGroupRedPackRequest 发裂变红包
type SendGroupRedPackRequest struct {
	XMLName     xml.Name `xml:"xml"`
	NonceStr    string   `xml:"nonce_str"`
	Sign        string   `xml:"sign"`
	MchBillno   string   `xml:"mch_billno"`
	MchID       string   `xml:"mch_id"`
	SendName    string   `xml:"send_name"`
	ReOpenID    string   `xml:"re_openid"`
	TotalAmount uint     `xml:"total_amount"`
	TotalNum    uint     `xml:"total_num"`
	AmtType     string   `xml:"amt_type"`
	Wishing     string   `xml:"wishing"`
	ActName     string   `xml:"act_name"`
	Remark      string   `xml:"remark"`
	SceneID     string   `xml:"scene_id"`
	RiskInfo    string   `xml:"risk_info"`
}

// NewSendGroupRedPackRequest 创建裂变红包
func NewSendGroupRedPackRequest() *SendGroupRedPackRequest {
	return &SendGroupRedPackRequest{
		NonceStr: random.AZaz09(32),
	}
}

// SignMD5 md5
func (req *SendGroupRedPackRequest) SignMD5(apiKey string) error {
	params, err := SignStructToParameter(*req)
	if err != nil {
		return err
	}
	params["key"] = apiKey
	value := SignMD5(params)
	req.Sign = value
	return nil
}

// SendGroupRedPackResponse ...
type SendGroupRedPackResponse struct {
	ResponseStatus
	MchBillno   string `xml:"mch_billno"`
	MchID       string `xml:"mch_id"`
	WxAppID     string `xml:"wxappid"`
	ReOpenID    string `xml:"re_openid"`
	TotalAmount uint   `xml:"total_amount"`
	SendListID  uint   `xml:"send_listid"`
}
