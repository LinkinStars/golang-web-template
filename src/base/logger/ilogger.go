package logger

type LogLevel int

const (
	logDebug LogLevel = iota
	logInfo
	logWarn
	logErr
	logOff
)

// 定义日志接口，如果后续需要替换日志框架只需要实现接口，替换实现即可
type ILogger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Level() LogLevel
	SetLevel(l LogLevel)
}
