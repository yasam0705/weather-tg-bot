package weather

type currentWeatherResponse struct {
	Location *location `json:"location"`
	Current  *current  `json:"current"`
}

type location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type current struct {
	LastUpdatedEpoch int        `json:"last_updated_epoch"`
	LastUpdated      string     `json:"last_updated"`
	TempC            float64    `json:"temp_c"`
	TempF            float64    `json:"temp_f"`
	IsDay            int        `json:"is_day"`
	Condition        *condition `json:"condition"`
	WindMph          float64    `json:"wind_mph"`
	WindKph          float64    `json:"wind_kph"`
	WindDegree       int        `json:"wind_degree"`
	WindDir          string     `json:"wind_dir"`
	PressureMb       float64    `json:"pressure_mb"`
	PressureIn       float64    `json:"pressure_in"`
	PrecipMm         float64    `json:"precip_mm"`
	PrecipIn         float64    `json:"precip_in"`
	Humidity         int        `json:"humidity"`
	Cloud            int        `json:"cloud"`
	FeelslikeC       float64    `json:"feelslike_c"`
	FeelslikeF       float64    `json:"feelslike_f"`
	VisKm            float64    `json:"vis_km"`
	VisMiles         float64    `json:"vis_miles"`
	Uv               float64    `json:"uv"`
	GustMph          float64    `json:"gust_mph"`
	GustKph          float64    `json:"gust_kph"`
}

type condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type currentError struct {
	Error *apiError `json:"error"`
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
