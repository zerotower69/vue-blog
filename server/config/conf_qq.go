package config

// QQ qq的配置
type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"`
}

// GetPath TODO:return the redirect url
func (q QQ) GetPath() string {
	return ``
}
