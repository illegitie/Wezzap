package models

type Forecast struct {
	Day     string             `json:"day"`
	Min     float64            `json:"temp_min"`
	Max     float64            `json:"temp_max"`
	Weather []WeatherCondition `json:"weather"`
}

type ForecastEvery3Hours struct {
	Time    string             `json:"day"`
	Temp    float64            `json:"temp"`
	Weather []WeatherCondition `json:"weather"`
}

type WeatherCondition struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type CurrentWeatherForecast struct {
	Day        string             `json:"day"`
	Temp       float64            `json:"temp"`
	FeelsLike  float64            `json:"feels_like"`
	Min        float64            `json:"temp_min"`
	Max        float64            `json:"temp_max"`
	Pressure   float64            `json:"pressure"`
	Humidity   int                `json:"humidity"`
	WindSpeed  float64            `json:"speed"`
	WindDeg    int                `json:"deg"`
	Clouds     int                `json:"clouds"`
	Visibility int                `json:"visibility"`
	Weather    []WeatherCondition `json:"weather"`
	Name       string             `json:"name"`
}
