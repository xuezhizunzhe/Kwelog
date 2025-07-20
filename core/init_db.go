package core

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"kweblog/global"
	"time"
)

func InitDB() *gorm.DB {
	if len(global.Config.DB) == 0 {
		logrus.Error("未配置数据库")
	}
	dc := global.Config.DB[0]
	db, err := gorm.Open(mysql.Open(dc.DNS()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败%s", err) // 触发Fatal方法会直接退出程序
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("数据库连接成功")

	if len(global.Config.DB) > 1 {
		var readList []gorm.Dialector
		for _, d := range global.Config.DB[1:] {
			readList = append(readList, mysql.Open(d.DNS()))
		}
		err = db.Use(dbresolver.Register(dbresolver.Config{
			// 写库有一个
			Sources: []gorm.Dialector{mysql.Open(dc.DNS())}, // 写
			// 读库有多个
			Replicas: readList, // 读 读库列表
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err != nil {
			logrus.Fatalf("读写配置错误 %s", err)
		}
	}
	return db
}
