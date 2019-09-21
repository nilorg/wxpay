package wxpay

const (
	Fail            = "FAIL"
	Success         = "SUCCESS"
	Ok              = "OK"
	HMACSHA256      = "HMAC-SHA256"
	MD5             = "MD5"
	TradeTypeJSApi  = "JSAPI"
	TradeTypeNative = "NATIVE"
	TradeTypeApp    = "APP"
)

var (
	CallbackSuccessful = &ResponseReturn{Code: Success, Msg: Ok}
)
