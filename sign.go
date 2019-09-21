package wxpay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"sort"
	"strings"

	"github.com/nilorg/sdk/convert"
)

func signBuffer(params Parameter) *bytes.Buffer {
	// 获取Key
	keys := []string{}
	for k := range params {
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	value := new(bytes.Buffer)
	for _, k := range keys {
		value.WriteString(k)
		value.WriteString("=")
		value.WriteString(interfaceToString(params[k]))
	}
	return value
}

// SignMD5 生成signMD5
func SignMD5(params Parameter) string {
	value := signBuffer(params)
	// 使用MD5加密
	h := md5.New()
	io.Copy(h, value)
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// SignHMACSHA256 生成Sign HMACS-HA256
func SignHMACSHA256(params Parameter, key []byte) string {
	value := signBuffer(params)
	// 使用HMACS-HA256加密
	h := hmac.New(sha256.New, key)
	io.Copy(h, value)
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func interfaceToString(src interface{}) string {
	if src == nil {
		panic(ErrTypeIsNil)
	}
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return convert.ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}
