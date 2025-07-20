package enum

type LogLevelType int8

const (
	LogInfoLevel LogLevelType = 1 // 登录日志
	LogWarnLevel LogLevelType = 2 // 操作日志
	LogErrLevel  LogLevelType = 3 // 运行日志
)

func (l LogLevelType) String() string {
	switch l {
	case LogInfoLevel:
		return "info"
	case LogWarnLevel:
		return "warn"
	case LogErrLevel:
		return "error"
	}
	return ""
}
