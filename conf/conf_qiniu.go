package conf

// 对象存储配置 --》 一定用的是七牛云
type Qiniu struct {
	Enable    bool   `yaml:"enable" json:"enable"`
	AccessKey string `yaml:"accessKey" json:"accessKey"`
	SecretKey string `yaml:"secretKey" json:"secretKey"`
	Bucket    string `yaml:"bucket" json:"bucket"` // 存储桶
	Uri       string `yaml:"uri" json:"uri"`       // 访问域名
	Region    string `yaml:"region" json:"region"` // 区域
	Prefix    string `yaml:"prefix" json:"prefix"` // 前缀
	Size      int    `yaml:"size" json:"size"`     // 单文件大小限制 (单位：MB)
	Expiry    int    `yaml:"expiry" json:"expiry"` // 链接有效期 (单位：秒)
}
