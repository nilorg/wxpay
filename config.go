package wxpay

// Config ...
type Config struct {
	BaseURL    string `json:"base_url"`
	AppID      string `json:"app_id"`
	MchID      string `json:"mch_id"`
	APIKey     string `json:"api_key"`
	SignType   string `json:"sign_type"`
	SignKey    string `json:"sign_key"`
	NotifyURL  string `json:"notify_url"`
	CaCertFile string `json:"ca_cert_file"`
	CaKeyFile  string `json:"ca_key_file"`
}

// NewConfig ...
func NewConfig(aid, mid string) *Config {
	return &Config{
		BaseURL:  BaseURL,
		AppID:    aid,
		MchID:    mid,
		SignType: "MD5",
	}
}
