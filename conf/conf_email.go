package conf

type Email struct {
	Domain       string `yaml:"domain" json:"domain"`             // 邮箱的域名
	Port         int    `yaml:"port" json:"port"`                 // 邮箱的端口
	SendEmail    string `yaml:"sendEmail" json:"sendEmail"`       // 发件人邮箱
	AuthCode     string `yaml:"authCode" json:"authCode"`         // 发件人邮箱的授权码
	SendNickname string `yaml:"sendNickname" json:"sendNickname"` // 发件人昵称
	SSL          bool   `yaml:"ssl" json:"ssl"`                   // 是否使用SSL加密
	TLS          bool   `yaml:"tls" json:"tls"`                   // 是否使用TLS加密

}
