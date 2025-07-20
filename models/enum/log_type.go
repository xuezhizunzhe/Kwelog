package enum

type LogType int8

const (
	LoginLogType   LogType = 1 // 登录日志
	ActionLogType  LogType = 2 // 操作日志
	RuntimeLogType LogType = 3 // 运行日志
)
