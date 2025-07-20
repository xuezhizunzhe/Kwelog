package conf

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     []DB   `yaml:"db"`
	Site   Site   `yaml:"site"`
	Jwt    Jwt    `yaml:"jwt"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	Qiniu  Qiniu  `yaml:"qiniu"`
	Ai     Ai     `yaml:"ai"`
}
