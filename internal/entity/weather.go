package entity

type CurrentWeather struct {
	Time          string
	Country       string
	Name          string
	TemperatureC  float64
	FeelslikeC    float64
	ConditionText string
	WindKph       float64
	Humidity      int
	Cloud         int
	GustKph       float64
}
