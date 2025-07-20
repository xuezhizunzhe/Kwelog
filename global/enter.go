package global

import (
	"gorm.io/gorm"
	"kweblog/conf"
)

const Version = "v10.0.1"

var (
	Config *conf.Config
	DB     *gorm.DB
)
