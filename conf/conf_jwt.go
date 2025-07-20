package conf

type Jwt struct {
	Expire int    `yaml:"expire"` // 过期时间
	Secret string `yaml:"secret"` // 密钥
	Issuer string `yaml:"issuer"` // 签发者
}
