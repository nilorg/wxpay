package wxpay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrTypeIsNil ...
	ErrTypeIsNil = errors.New("类型为Nil")
	// ErrTypeUnknown ...
	ErrTypeUnknown = errors.New("未处理到的数据类型")
)

var (
	// Error ...
	resultErrors = map[string]string{
		"INVALID_REQUEST":       "参数错误",
		"NOAUTH":                "商户无此接口权限",
		"NOTENOUGH":             "余额不足",
		"ORDERPAID":             "商户订单已支付",
		"ORDERCLOSED":           "订单已关闭",
		"SYSTEMERROR":           "系统错误",
		"APPID_NOT_EXIST":       "APPID不存在",
		"MCHID_NOT_EXIST":       "MCHID不存在",
		"APPID_MCHID_NOT_MATCH": "appid和mch_id不匹配",
		"LACK_PARAMS":           "缺少参数",
		"OUT_TRADE_NO_USED":     "商户订单号重复",
		"SIGNERROR":             "签名错误",
		"XML_FORMAT_ERROR":      "XML格式错误",
		"REQUIRE_POST_METHOD":   "请使用post方法",
		"POST_DATA_EMPTY":       "post数据为空",
		"NOT_UTF8":              "编码格式错误",
	}
)

// GetErrorMsg 获取error msg
func GetErrorMsg(code string) string {
	if msg, ok := resultErrors[code]; ok {
		return msg
	}
	return fmt.Sprintf("错误Code：%s未知", code)
}

// errorModel 错误模型
type errorModel struct {
	Code string `xml:"err_code"`
	Msg  string `xml:"err_code_des"`
}

// NewError 创建错误
func NewError(buf []byte) error {
	errstr := string(buf)
	if strings.Index(errstr, `{"err_code":`) == -1 && strings.Index(errstr, `,"err_code_des":"`) == -1 {
		return nil
	}
	model := new(errorModel)
	err := xml.Unmarshal(buf, model)
	if err != nil {
		return nil
	}
	if model.Code == "" {
		return nil
	}
	return errors.New(GetErrorMsg(model.Code))
}
