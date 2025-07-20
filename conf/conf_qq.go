package conf

import "fmt"

type QQ struct {
	AppID    string `yaml:"appID" json:"appID"`
	AppKey   string `yaml:"appKey" json:"appKey"`
	Redirect string `yaml:"redirect" json:"redirect"`
}

func (q QQ) Url() string { // 回调地址  ---> 只需要有 AppID 和 Redirect(回调地址)就可以拼出来 URL
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
