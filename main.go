package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	weatherforecast "wf/Forecast"
)

func main() {
	resp, err := http.Get("https://www.meteosource.com/api/v1/free/point?lat=50.4501&lon=30.5234&sections=daily&language=en&units=auto&key=3wdq2a8ni3mmvvdsahki3w38l5g3hmqbn5jclpiy")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data weatherforecast.DailyData
	json.Unmarshal(jsonBody, &data)

	for _, f := range data.Daily.Forecast {
		f.PrintData()
		println("\n")
	}
	
	f, err := data.Daily.GetForecast("2024-12-26")
	f.PrintData()
}
