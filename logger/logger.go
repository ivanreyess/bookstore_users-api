package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	//Log entry point for logger library
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

//Info logs the message with string message and different tags
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

//Error logs the error message with string message and different tags
func Error(msg string, tags ...zap.Field) {
	log.Error(msg, tags...)
	log.Sync()
}

//GetLogger return zap logger instance
func GetLogger() *zap.Logger {
	return log
}
