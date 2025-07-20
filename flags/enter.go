package flags

import (
	"flag"
	"os"
)

type Options struct {
	File string
	DB   bool
}

var FlagOptions = new(Options)

func Parse() {
	// flag
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "b", false, "数据库迁移")
	flag.Parse()
}

func Run() {
	// 执行数据库迁移
	if FlagOptions.DB {
		// 迁移数据库
		FlagDB()
		// 如果用户加了 --db 参数，只做迁移，做完就“收工”，不继续往下走业务逻辑。
		os.Exit(0)
	}
}
