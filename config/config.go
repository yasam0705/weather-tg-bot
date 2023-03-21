package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	APP              string
	Environment      string
	CtxTimeout       time.Duration
	LogLevel         string
	TelegramBotToken string
	Postgres         struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
	Weather struct {
		Scheme  string
		BaseUrl string
		ApiKey  string
	}
}

func New() (*Config, error) {
	config := &Config{}

	config.APP = getEnv("APP", "test-tg-bot")
	config.LogLevel = getEnv("LOG_LEVEL", "debug")
	ctxTimeout, err := time.ParseDuration(getEnv("CONTEXT_TIMEOUT", "7s"))
	if err != nil {
		return nil, fmt.Errorf("error on parse duration: %s", err.Error())
	}
	config.CtxTimeout = ctxTimeout

	config.Environment = getEnv("ENVIRONMENT", "dev")
	config.TelegramBotToken = getEnv("BOT_TOKEN", "token")

	config.Postgres.Host = getEnv("POSTGRES_HOST", "localhost")
	config.Postgres.Port = getEnv("POSTGRES_PORT", "5432")
	config.Postgres.User = getEnv("POSTGRES_USER", "sam")
	config.Postgres.Password = getEnv("POSTGRES_PASSWORD", "")
	config.Postgres.Database = getEnv("POSTGRES_DATABASE", "telegram_bot")

	config.Weather.Scheme = getEnv("WEATHER_SCHEME", "https")
	config.Weather.BaseUrl = getEnv("WEATHER_BASE_URL", "api.weatherapi.com")
	config.Weather.ApiKey = getEnv("WEATHER_API_KEY", "key")

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
