package wxpay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const bodyType = "application/xml; charset=utf-8"

// Clienter ...
type Clienter interface {
	Config() Config

	UnifiedOrder(req *UnifiedOrderRequest) (resp *UnifiedOrderResponse, err error)
}

// Client 客户端
type Client struct {
	conf *Config
}

// NewClient ...
func NewClient(conf *Config) *Client {
	return &Client{
		conf: conf,
	}
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

	resp, err := http.Post(c.conf.BaseURL+url, bodyType, value)
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
