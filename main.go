package main

import (
	"kweblog/core"
	"kweblog/flags"
	"kweblog/global"
)

func main() {
	flags.Parse() // 解析命令行参数

	global.Config = core.ReadConf() // 读取配置
	core.InitLogrus()               // 初始化日志
	global.DB = core.InitDB()       // 初始化数据库
	flags.Run()                     // 执行命令行参数
}
