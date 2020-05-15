package main

import (
	"encoding/json"
	"time"
)

type DarkSky_Forecast struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly DarkSky_Report `json:"hourly"`
}

type DarkSky_Report struct {
	Summary string `json:"summary"`
	Icon string `json:"icon"`
	Data []DarkSky_ReportData `json:"data"`
}

type DarkSky_ReportData struct {
	Time time.Time `json:"time"`
	Temperature float64 `json:"temperature"`
	WindSpeed float64 `json:"windSpeed"`
	WindBearing float64 `json:"windBearing"`
	WindGust float64 `json:"windGust"`
	ApparentTemperature float64 `json:"apparentTemperature"`
	Humidity float64 `json:"humidity"`
	UvIndex float64 `json:"uvIndex"`
	CloudCover float64 `json:"cloudCover"`
	PrecipIntensity float64 `json:"precipIntensity"`
	Pressure float64 `json:"pressure"`
	Icon string `json:"icon"`
}

func (s *DarkSky_ReportData) UnmarshalJSON(data []byte) error {
	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil { return err }
	m := f.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "temperature": s.Temperature  = v.(float64)
		case "windSpeed": s.WindSpeed = v.(float64)
		case "windBearing": s.WindBearing  = v.(float64)
		case "windGust": s.WindGust  = v.(float64)
		case "apparentTemperature": s.ApparentTemperature  = v.(float64)
		case "humidity": s.Humidity  = v.(float64)
		case "uvIndex": s.UvIndex  = v.(float64)
		case "cloudCover": s.CloudCover  = v.(float64)
		case "precipIntensity": s.PrecipIntensity  = v.(float64)
		case "pressure": s.Pressure  = v.(float64)
		case "icon": s.Icon  = v.(string)
		case "time": s.Time = time.Unix(int64(v.(float64)), 0)
		}
	}

	return nil
}