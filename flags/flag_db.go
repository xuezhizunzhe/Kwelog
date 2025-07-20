package flags

import (
	"github.com/sirupsen/logrus"
	"kweblog/global"
	"kweblog/models"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.LogModel{},
	)
	if err != nil {
		logrus.Errorf("数据库迁移失败%s", err)
		return
	}
	logrus.Infof("数据库迁移成功")
}
