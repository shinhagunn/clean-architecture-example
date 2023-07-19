package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
)

func Setup() {
	config := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevel(),
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			LevelKey:     "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeLevel:  CustomLevelEncoder,
			EncodeTime:   SyslogTimeEncoder,
		},
	}

	log, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger = log.Sugar()
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func Debug(skip int, template string, args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(skip)).Debugf(template, args...)
}

func Info(skip int, template string, args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(skip)).Infof(template, args...)
}

func Warn(skip int, template string, args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(skip)).Warnf(template, args...)
}

func Error(skip int, template string, args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(skip)).Errorf(template, args...)
}

func Fatal(skip int, template string, args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(skip)).Fatalf(template, args...)
}
