package logger

// ILogger 定义日志接口，如果后续需要替换日志框架只需要实现接口，替换实现即可
type ILogger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

// Debug log
func Debug(v ...interface{}) {
	myLogger.Debug(v...)
}

// Debugf log
func Debugf(format string, v ...interface{}) {
	myLogger.Debugf(format, v...)
}

// Info log
func Info(v ...interface{}) {
	myLogger.Info(v...)
}

// Infof log
func Infof(format string, v ...interface{}) {
	myLogger.Infof(format, v...)
}

// Warn log
func Warn(v ...interface{}) {
	myLogger.Warn(v...)
}

// Warnf log
func Warnf(format string, v ...interface{}) {
	myLogger.Warnf(format, v...)
}

// Error log
func Error(v ...interface{}) {
	myLogger.Error(v...)
}

// Errorf log
func Errorf(format string, v ...interface{}) {
	myLogger.Errorf(format, v...)
}
