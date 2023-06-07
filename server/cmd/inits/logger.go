package inits

import (
	"os"
	"path"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/bagasunix/kuyngetrip/pkg/helpers"
)

var (
	lock    = &sync.Mutex{}
	Loggers *zap.Logger
)

func InitLogger() *zap.Logger {
	writerSyncer := getLogWriter()
	var core zapcore.Core

	if os.Getenv("ENV") == "dev" {
		core = zapcore.NewTee(
			zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(getFileEncoder(), writerSyncer, zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(getFileEncoder(), writerSyncer, zapcore.DebugLevel)
	}

	Loggers = zap.New(core, zap.AddCaller())
	defer Loggers.Sync() // flushes buffer, if any

	return Loggers
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.StacktraceKey = zapcore.DPanicLevel.String()
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {

	logFileName := "kuyngetrip -" + helpers.DateFormatEN() + ".log"
	logFile := path.Join("../../", logFileName)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
