package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var logger = logrus.New()

func Initial(level logrus.Level, logPath string) {
	formatter := &Formatter{
		LogFormat:       "%time% [%lvl%] %msg%",
		TimestampFormat: "2006-01-02 15:04:05",
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetFormatter(formatter)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(level)

	// Output to file
	logFilePath := filepath.Join(logPath, "./log/new.log")
	rotateFileHook, err := NewRotateFileHook(RotateFileConfig{
		Filename:   logFilePath,
		MaxSize:    50,
		MaxBackups: 1024,
		MaxAge:     30,
		LocalTime:  true,
		Level:      level,
		Formatter:  formatter,
	})
	if err != nil {
		fmt.Printf("Create log rotate hooks error: %s\n", err)
		return
	}
	logger.AddHook(rotateFileHook)
}

func Debugln(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infoln(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnln(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorln(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
