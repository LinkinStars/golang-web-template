package logger

import (
	"time"

	"go.uber.org/zap"
)

// myLogger 日志内部实现
var myLogger *zapLogger

// levelMapping 日志级别的映射关系
var levelMapping = levelMap()

// logLevel 日志级别
type logLevel int

const (
	logDebug logLevel = iota
	logInfo
	logWarn
	logErr
	logOff
)

// InitLogger 初始化日志Logger
func InitLogger(level, projectName, logPath string, maxAge, rotationTime time.Duration) {
	initZap(projectName, logPath, maxAge, rotationTime)
	// 设置打印堆栈深度，设置日志级别
	myLogger = &zapLogger{
		slog:  zap.L().WithOptions(zap.AddCallerSkip(2)).Sugar(),
		level: levelMapping[level],
	}
}

// levelMap 字符串和级别映射
func levelMap() map[string]logLevel {
	return map[string]logLevel{
		"debug": logDebug,
		"info":  logInfo,
		"warn":  logWarn,
		"err":   logErr,
		"off":   logOff,
	}
}

// zapLogger 具体对内实现
type zapLogger struct {
	slog  *zap.SugaredLogger
	level logLevel
}

// Debug log
func (z *zapLogger) Debug(v ...interface{}) {
	if z.level <= logDebug {
		z.slog.Debug(v...)
	}
}

// Debugf log
func (z *zapLogger) Debugf(format string, v ...interface{}) {
	if z.level <= logDebug {
		z.slog.Debugf(format, v...)
	}
}

// Info log
func (z *zapLogger) Info(v ...interface{}) {
	if z.level <= logInfo {
		z.slog.Info(v...)
	}
}

// Infof log
func (z *zapLogger) Infof(format string, v ...interface{}) {
	if z.level <= logInfo {
		z.slog.Infof(format, v...)
	}
}

// Warn log
func (z *zapLogger) Warn(v ...interface{}) {
	if z.level <= logWarn {
		z.slog.Warn(v...)
	}
}

// Warnf log
func (z *zapLogger) Warnf(format string, v ...interface{}) {
	if z.level <= logWarn {
		z.slog.Warnf(format, v...)
	}
}

// Error log
func (z *zapLogger) Error(v ...interface{}) {
	if z.level <= logErr {
		z.slog.Error(v...)
	}
}

// Errorf log
func (z *zapLogger) Errorf(format string, v ...interface{}) {
	if z.level <= logErr {
		z.slog.Errorf(format, v)
	}
}
