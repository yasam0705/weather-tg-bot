package logger

import (
	"test-tasks/tg-bot/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg *config.Config) (*zap.Logger, error) {
	var log zap.Config

	if cfg.Environment == config.ProdEnvironment {
		// prod logger
	} else {
		log = newDevLogger(cfg)
	}

	return log.Build()
}

func newDevLogger(cfg *config.Config) zap.Config {
	zapCfg := zap.NewDevelopmentConfig()

	level := getLevel(cfg.LogLevel)

	zapCfg.Level = zap.NewAtomicLevelAt(level)
	return zapCfg
}

func getLevel(l string) zapcore.Level {
	switch l {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.DebugLevel
	}
}
