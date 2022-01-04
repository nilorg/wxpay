package wxpay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"reflect"
	"sort"
	"strings"

	"github.com/nilorg/sdk/convert"
)

func signBuffer(params Parameter) *bytes.Buffer {
	// 获取Key
	keys := []string{}
	for k := range params {
		if k == "key" {
			continue
		}
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
		value.WriteString("&")
	}
	value.WriteString("key=")
	value.WriteString(convert.ToString(params["key"]))
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
	switch v := src.(type) {
	case string:
		return v
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return convert.ToString(v)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// SignStructToParameter ...
func SignStructToParameter(value interface{}) (params Parameter, err error) {
	params = Parameter{}
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			xname := f.Tag.Get("xml")
			if xname == "-" || xname == "xml" || xname == "sign" {
				continue
			}
			xvalue := v.FieldByName(f.Name).Interface()
			if convert.ToString(xvalue) != "" {
				params[xname] = xvalue
			}
		}
	default:
		err = ErrNotEqualStruct
	}
	return
}
