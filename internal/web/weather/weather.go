package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"test-tasks/tg-bot/internal/entity"
	"time"
)

type weather struct {
	scheme  string
	baseUrl string
	apiKey  string
	client  *http.Client
}

type WeatherService interface {
	CurrentWeather(ctx context.Context, city string) (*entity.CurrentWeather, error)
}

func New(scheme, baseUrl, apiKey string) *weather {
	return &weather{
		scheme:  scheme,
		baseUrl: baseUrl,
		apiKey:  apiKey,
		client:  &http.Client{},
	}
}

func (w *weather) CurrentWeather(ctx context.Context, city string) (*entity.CurrentWeather, error) {
	params := url.Values{
		"key":  []string{w.apiKey},
		"q":    []string{city},
		"aqi":  []string{"no"},
		"lang": []string{"ru"},
	}
	u := url.URL{
		Scheme:   w.scheme,
		Host:     w.baseUrl,
		Path:     "/v1/current.json",
		RawQuery: params.Encode(),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	req.Close = true
	// response, err := http.Get(url)

	response, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = w.handleError(response)
		return nil, err
	}

	weather := new(currentWeatherResponse)
	if err = json.NewDecoder(response.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &entity.CurrentWeather{
		Time:          time.Unix(weather.Location.LocaltimeEpoch, 0).Format("2006-01-02 15:04"),
		Country:       weather.Location.Country,
		Name:          weather.Location.Name,
		TemperatureC:  weather.Current.TempC,
		FeelslikeC:    weather.Current.FeelslikeC,
		ConditionText: weather.Current.Condition.Text,
		WindKph:       weather.Current.WindKph,
		Humidity:      weather.Current.Humidity,
		Cloud:         weather.Current.Cloud,
		GustKph:       weather.Current.GustKph,
	}, nil
}

func (w *weather) handleError(r *http.Response) error {
	var (
		msg        string
		weatherErr = new(currentError)
	)

	if err := json.NewDecoder(r.Body).Decode(&weatherErr); err != nil {
		return err
	}

	switch weatherErr.Error.Code {
	case 1006:
		msg = "No location found"
	case 9999:
		msg = "External application error"
	default:
		msg = weatherErr.Error.Message
	}

	return fmt.Errorf(msg)
}
