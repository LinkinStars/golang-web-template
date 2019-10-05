// Copyright 2019 LinkinStar
// license that can be found in the LICENSE file.

package gu

import (
	"os"
	"path"
	"strings"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const defaultLogPath = "logs"
const defaultMaxAge = 7 * 24 * time.Hour
const defaultRotationTime = 24 * time.Hour

var yourProjectName = "golang-util"

// projectName : outermost catalogue of your project
func InitEasyZapDefault(projectName string) {
	InitEasyZap(projectName, defaultLogPath, defaultMaxAge, defaultRotationTime)
}

// logPath : the catalogue for output log;
// maxAge : maximum storage time of log files;
// rotationTime : how long a new log file will be generated;
// projectName : outermost catalogue of your project;
func InitEasyZap(projectName, logPath string, maxAge, rotationTime time.Duration) {
	if len(projectName) != 0 {
		yourProjectName = projectName
	}
	if err := CreateDirIfNotExist(logPath); err != nil {
		panic(err)
	}
	logPath = path.Join(logPath, projectName)

	// err log file
	errWriter, err := rotatelogs.New(
		logPath+"_err_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic(err)
	}

	// info log file
	infoWriter, err := rotatelogs.New(
		logPath+"_info_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic(err)
	}

	// Priority
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.InfoLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})

	// console output
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderConfig.EncodeTime = timeEncoder
	consoleEncoderConfig.EncodeCaller = customCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// file output
	errorCore := zapcore.AddSync(errWriter)
	infoCore := zapcore.AddSync(infoWriter)
	fileEncodeConfig := zap.NewProductionEncoderConfig()
	fileEncodeConfig.EncodeTime = timeEncoder
	fileEncodeConfig.EncodeCaller = customCallerEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileEncodeConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, errorCore, highPriority),
		zapcore.NewCore(fileEncoder, infoCore, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, zapcore.DebugLevel),
	)

	// show line number
	caller := zap.AddCaller()
	// you can set caller skip to skip some stack logs. mind that add it to zap.New option
	// callerSkip := zap.AddCallerSkip(1)

	development := zap.Development()
	logger := zap.New(core, caller, development)

	// replace global logger
	zap.ReplaceGlobals(logger)

	// redirect std log to zap log and log level is error
	if _, err := zap.RedirectStdLogAt(logger, zapcore.ErrorLevel); err != nil {
		panic(err)
	}
}

// custom caller to make the log can output the full path from "src"
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	str := caller.String()
	index := strings.Index(str, yourProjectName)
	if index == -1 {
		enc.AppendString(caller.FullPath())
	} else {
		index = index + len(yourProjectName) + 1
		enc.AppendString(str[index:])
	}
}

// Formatting log time
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
