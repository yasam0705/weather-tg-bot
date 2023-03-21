package app

import (
	"context"
	cfg_pkg "test-tasks/tg-bot/config"
	"test-tasks/tg-bot/internal/delivery/telegram"
	"test-tasks/tg-bot/internal/repository/postgres"
	"test-tasks/tg-bot/internal/usecase"
	"test-tasks/tg-bot/internal/web/weather"

	postgres_pkg "test-tasks/tg-bot/pkg/postgres"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func Run(config *cfg_pkg.Config, log *zap.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	client, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Error("error on tgbotapi.NewBotAPI", zap.Error(err))
		return err
	}

	db, err := postgres_pkg.New(ctx, config)
	if err != nil {
		log.Error("error on postgres_pkg.New", zap.Error(err))
		return err
	}

	clientRepo := postgres.NewClientRepo(db)
	messageRepo := postgres.NewMessageRepo(db)

	weatherService := weather.New(config.Weather.Scheme, config.Weather.BaseUrl, config.Weather.ApiKey)

	clientUseCase := usecase.NewClient(config.CtxTimeout, clientRepo)
	messageUseCase := usecase.NewMessage(config.CtxTimeout, messageRepo)
	statUseCase := usecase.NewStat(messageUseCase, config.CtxTimeout)
	weatherUseCase := usecase.NewWeather(config.CtxTimeout, weatherService)

	bot := telegram.New(client, clientUseCase, messageUseCase, log, statUseCase, config.CtxTimeout, weatherUseCase)

	bot.Run()
	return nil
}
