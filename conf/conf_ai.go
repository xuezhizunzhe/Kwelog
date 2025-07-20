package conf

// AI 设置
// 用于系统中的文件推荐功能，若不配置则不启用
type Ai struct {
	Enable    bool   `yaml:"enable" json:"enable"`       // 是否开启AI功能
	SecretKey string `yaml:"secretKey" json:"secretKey"` // 密钥
	Nickname  string `yaml:"nickname" json:"nickname"`   // 昵称
	Avatar    string `yaml:"avatar" json:"avatar"`       // 头像
	Abstract  string `yaml:"abstract" json:"abstract"`   // 简介
}
