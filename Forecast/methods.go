package weatherforecast

import (
	"errors"
	"fmt"
)

// date example: 2024-12-24
func (d *Daily) GetForecast(date string) (*Forecast, error) {
	for _, f := range d.Forecast {
		if f.Day == date {
			return &f, nil
		}
	}
	return nil, errors.New("date error")
}

func (f *Forecast) PrintData() {
	fmt.Printf("Date: %s\nSummary: %s\nMax temperature: %f\nMin temperature: %f\nWeather: %s",
		f.Day, f.Summary, f.TemperatureMax, f.TemperatureMin, f.Weather,
	)
}
