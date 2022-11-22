package wxpay

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const bodyType = "application/xml; charset=utf-8"

// Clienter ...
type Clienter interface {
	Config() Config

	UnifiedOrder(req *UnifiedOrderRequest) (resp *UnifiedOrderResponse, err error)
}

// Client 客户端
type Client struct {
	conf       *Config
	httpClient *http.Client
}

// NewClient ...
func NewClient(conf *Config) (client *Client, err error) {
	var transport http.RoundTripper
	if conf.CaCertFile != "" && conf.CaKeyFile != "" {
		var cert tls.Certificate
		cert, err = tls.LoadX509KeyPair(conf.CaCertFile, conf.CaKeyFile)
		if err != nil {
			return
		}
		t := &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		}
		if conf.HttpProxy != "" {
			var proxyURL *url.URL
			proxyURL, err = url.Parse(conf.HttpProxy)
			if err != nil {
				return
			}
			t.Proxy = http.ProxyURL(proxyURL)
		}
		transport = t
	} else {
		defaultTransport := http.DefaultTransport
		if conf.HttpProxy != "" {
			var proxyURL *url.URL
			proxyURL, err = url.Parse(conf.HttpProxy)
			if err != nil {
				return
			}
			defaultTransport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
		}
		transport = defaultTransport
	}
	client = &Client{
		conf: conf,
		httpClient: &http.Client{
			Transport: transport,
		},
	}
	return
}

// Config ...
func (c *Client) Config() Config {
	return *c.conf
}

// ResponseStatus 响应状态
type ResponseStatus struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

// ResponseReturn 成功
type ResponseReturn struct {
	XMLName xml.Name `xml:"xml"`
	Code    CDATA    `xml:"return_code"`
	Msg     CDATA    `xml:"return_msg"`
}

// execute 执行
func (c *Client) execute(url string, param interface{}) (body []byte, err error) {
	value := new(bytes.Buffer)
	xen := xml.NewEncoder(value)
	err = xen.Encode(param)
	if err != nil {
		return
	}
	resp, err := c.httpClient.Post(c.conf.BaseURL+url, bodyType, value)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

// UnifiedOrder 统一下单
func (c *Client) UnifiedOrder(req *UnifiedOrderRequest) (resp *UnifiedOrderResponse, err error) {
	req.AppID = c.conf.AppID
	req.MchID = c.conf.MchID
	if req.NotifyURL == "" {
		req.NotifyURL = c.conf.NotifyURL
	}

	if c.conf.SignType == HMACSHA256 {
		err = req.SignHMACSHA256([]byte(c.conf.SignKey), c.conf.APIKey)
		if err != nil {
			return
		}
	} else {
		err = req.SignMD5(c.conf.APIKey)
		if err != nil {
			return
		}
	}

	var body []byte
	body, err = c.execute("/pay/unifiedorder", req)
	if err != nil {
		return
	}

	resp = new(UnifiedOrderResponse)
	err = xml.Unmarshal(body, resp)
	if err != nil {
		resp = nil
		return
	}
	if resp.ReturnCode == "FAIL" {
		err = fmt.Errorf("通信错误：%s", resp.ReturnMsg)
		return
	}
	if resp.ResultCode == "FAIL" {
		err = fmt.Errorf("业务错误：%s", resp.ErrCodeDes)
		return
	}
	return
}

// SendRedPack 发现金红包
func (c *Client) SendRedPack(req *SendRedPackRequest) (resp *SendRedPackResponse, err error) {
	req.WxAppID = c.conf.AppID
	req.MchID = c.conf.MchID
	err = req.SignMD5(c.conf.APIKey)
	if err != nil {
		return
	}

	var body []byte
	body, err = c.execute("/mmpaymkttransfers/sendredpack", req)
	if err != nil {
		return
	}

	resp = new(SendRedPackResponse)
	err = xml.Unmarshal(body, resp)
	if err != nil {
		resp = nil
		return
	}
	if resp.ReturnCode == "FAIL" {
		err = fmt.Errorf("通信错误：%s", resp.ReturnMsg)
		return
	}
	if resp.ResultCode == "FAIL" {
		err = fmt.Errorf("业务错误：%s", resp.ErrCodeDes)
		return
	}
	return
}

// SendGroupRedPack 发裂变红包
func (c *Client) SendGroupRedPack(req *SendGroupRedPackRequest) (resp *SendGroupRedPackResponse, err error) {
	req.WxAppID = c.conf.AppID
	req.MchID = c.conf.MchID
	err = req.SignMD5(c.conf.APIKey)
	if err != nil {
		return
	}

	var body []byte
	body, err = c.execute("/mmpaymkttransfers/sendgroupredpack", req)
	if err != nil {
		return
	}

	resp = new(SendGroupRedPackResponse)
	err = xml.Unmarshal(body, resp)
	if err != nil {
		resp = nil
		return
	}
	if resp.ReturnCode == "FAIL" {
		err = fmt.Errorf("通信错误：%s", resp.ReturnMsg)
		return
	}
	if resp.ResultCode == "FAIL" {
		err = fmt.Errorf("业务错误：%s", resp.ErrCodeDes)
		return
	}
	return
}

// PromotionTransfers 企业向微信用户个人付款
func (c *Client) PromotionTransfers(req *PromotionTransfersRequest) (resp *PromotionTransfersResponse, err error) {
	req.MchAppID = c.conf.AppID
	req.MchID = c.conf.MchID
	err = req.SignMD5(c.conf.APIKey)
	if err != nil {
		return
	}

	var body []byte
	body, err = c.execute("/mmpaymkttransfers/promotion/transfers", req)
	if err != nil {
		return
	}

	resp = new(PromotionTransfersResponse)
	err = xml.Unmarshal(body, resp)
	if err != nil {
		resp = nil
		return
	}
	if resp.ReturnCode == "FAIL" {
		err = fmt.Errorf("通信错误：%s", resp.ReturnMsg)
		return
	}
	if resp.ResultCode == "FAIL" {
		err = fmt.Errorf("业务错误：%s", resp.ErrCodeDes)
		return
	}
	return
}
