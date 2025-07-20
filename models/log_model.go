package models

import (
	"kweblog/models/enum"
)

type LogModel struct {
	Model
	LogType enum.LogType      `json:"logType"`              // 日志类型 1:登录日志 2:操作日志 3:运行日志
	Title   string            `gorm:"size:64" json:"title"` // 日志标题
	Content string            `json:"content"`              // 日志内容
	Level   enum.LogLevelType `json:"level"`                // 日志级别 1:一般日志 2:警告日志 3:错误日志
	UserID  uint              `json:"userID"`               // 用户ID
	//UserModel   UserModel         `gorm:"foreignKey:UserID" json:"-"` // 用户信息
	IP          string `gorm:"size:32" json:"ip"`       // IP地址
	Addr        string `gorm:"size:64" json:"addr"`     // 地址
	IsRead      bool   `json:"isRead"`                  // 是否已读
	LoginStatus bool   `json:"loginStatus"`             // 登录状态
	Username    string `gorm:"size:32" json:"username"` // 登录日志的用户名'
	Pwd         string `gorm:"size:32" json:"pwd"`      // 登录日志的密码
	//LoginType   enum.LoginType    `json:"loginType"`                  // 登录类型 1:密码登录 2:验证码登录 3:微信登录 4:QQ登录 5:微博登录
	ServiceName string `gorm:"size:32" json:"serviceName"` // 服务名称
}
