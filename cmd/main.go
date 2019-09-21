package main

import (
	"log"

	"github.com/nilorg/wxpay"
)

func main() {
	conf := wxpay.NewConfig("xxxx", "11112222")
	conf.APIKey = "123"
	// conf.APIKey = "123"
	conf.NotifyURL = "http://www.weixin.qq.com/wxpay/pay.php"
	conf.SignType = wxpay.MD5

	client := wxpay.NewClient(conf)

	uoReq := wxpay.NewUnifiedOrderRequest()
	uoReq.DeviceInfo = "WEB"
	uoReq.Body = "我惠淘-测试接口"
	// uoReq.Detail = wxpay.CDATA(`{ "goods_detail":[ { "goods_id":"iphone6s_16G", "wxpay_goods_id":"1001", "goods_name":"iPhone6s 16G", "quantity":1, "price":528800, "goods_category":"123456", "body":"苹果手机" }, { "goods_id":"iphone6s_32G", "wxpay_goods_id":"1002", "goods_name":"iPhone6s 32G", "quantity":1, "price":608800, "goods_category":"123789", "body":"苹果手机" } ] }`)
	uoReq.OutTradeNo = "20150806125346"
	uoReq.TotalFee = 100
	// 可自动处理
	uoReq.SpbillCreateIP = "123.12.12.123"
	uoReq.TradeType = wxpay.TradeTypeJSApi
	uoReq.OpenID = "oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"

	uoResp, err := client.UnifiedOrder(uoReq)
	if err != nil {
		log.Printf("统一下单错误:%s\n", err)
		return
	}
	log.Printf("结果：%v\n", uoResp)
}
