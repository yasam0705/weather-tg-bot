package weather_test

import (
	"context"
	"test-tasks/tg-bot/internal/web/weather"
	"testing"
)

func TestCurrentWeather(t *testing.T) {
	w := weather.New("https", "api.weatherapi.com", "key")

	resp, err := w.CurrentWeather(context.Background(), "Ташкент")
	if err != nil {
		t.Error(err)
	}
	_ = resp
}
