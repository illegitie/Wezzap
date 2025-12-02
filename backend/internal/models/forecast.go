package models

type Forecast struct {
	Day string  `json:"day"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
