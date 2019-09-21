package wxpay

import "errors"

var (
	// ErrTypeIsNil ...
	ErrTypeIsNil = errors.New("类型为Nil")
	// ErrTypeUnknown ...
	ErrTypeUnknown = errors.New("未处理到的数据类型")
)
