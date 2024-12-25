package weatherforecast

type DailyData struct {
	Daily Daily `json:"daily"`
}

type Daily struct {
	Forecast []Forecast `json:"data"`
}

type Forecast struct {
	Day            string  `json:"day"`
	Weather        string  `json:"weather"`
	Summary        string  `json:"summary"`
	TemperatureMin float64 `json:"temperature_min"`
	TemperatureMax float64 `json:"temperature_max"`
}
