package wxpay

import (
	"errors"
)

var (
	// ErrNotEqualStruct ...
	ErrNotEqualStruct = errors.New("参数不是Struct类型")
)

// Parameter 参数
type Parameter map[string]interface{}
