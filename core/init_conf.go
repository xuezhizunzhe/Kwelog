package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"kweblog/conf"
	"kweblog/flags"
	"os"
)

func ReadConf() (c *conf.Config) {
	bytedata, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	// 在内存分配地址
	c = new(conf.Config)
	err = yaml.Unmarshal(bytedata, c)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件格式错误%s", err))
	}
	fmt.Printf("读取配置文件%s成功\n", flags.FlagOptions.File) // 这一行输出结果 ： 读取配置文件settings.yaml成功
	return c
}
