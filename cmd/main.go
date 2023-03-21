package main

import (
	"test-tasks/tg-bot/config"
	"test-tasks/tg-bot/internal/app"
	log_pkg "test-tasks/tg-bot/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logger, err := log_pkg.New(cfg)
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Debug(cfg.APP + " start...")
	if err := app.Run(cfg, logger); err != nil {
		logger.Error("error app.Run", zap.Error(err))
	}

	logger.Debug(cfg.APP + " stop")
}
