package usecase

import (
	"context"
	"test-tasks/tg-bot/internal/entity"
	weather_web "test-tasks/tg-bot/internal/web/weather"
	"time"
)

type Weather interface {
	CurrentWeather(ctx context.Context, city string) (*entity.CurrentWeather, error)
}

type weather struct {
	ctxTimeout     time.Duration
	weatherService weather_web.WeatherService
}

func NewWeather(ctxTimeout time.Duration, weatherService weather_web.WeatherService) *weather {
	return &weather{
		ctxTimeout:     ctxTimeout,
		weatherService: weatherService,
	}
}

func (w *weather) CurrentWeather(ctx context.Context, city string) (*entity.CurrentWeather, error) {
	ctx, cancel := context.WithTimeout(ctx, w.ctxTimeout)
	defer cancel()

	return w.weatherService.CurrentWeather(ctx, city)
}
