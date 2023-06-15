package main

import (
	"encoding/json"
	"time"
)

type Tomorrow_Forecast struct {
	Timelines struct {
		Hourly []Tomorrow_Report `json:"hourly"`
	} `json:"timelines"`
}

type Tomorrow_Report struct {
	Time   time.Time       `json:"time"`
	Values Tomorrow_Values `json:"values"`
}

type Tomorrow_Values struct {
	CloudBase                float64     `json:"cloudBase"`
	CloudCeiling             interface{} `json:"cloudCeiling"`
	CloudCover               int         `json:"cloudCover"`
	DewPoint                 float64     `json:"dewPoint"`
	Evapotranspiration       float64     `json:"evapotranspiration"`
	FreezingRainIntensity    int         `json:"freezingRainIntensity"`
	Humidity                 int         `json:"humidity"`
	IceAccumulation          int         `json:"iceAccumulation"`
	IceAccumulationLwe       int         `json:"iceAccumulationLwe"`
	PrecipitationProbability int         `json:"precipitationProbability"`
	PressureSurfaceLevel     float64     `json:"pressureSurfaceLevel"`
	RainAccumulation         int         `json:"rainAccumulation"`
	RainAccumulationLwe      int         `json:"rainAccumulationLwe"`
	RainIntensity            int         `json:"rainIntensity"`
	SleetAccumulation        int         `json:"sleetAccumulation"`
	SleetAccumulationLwe     int         `json:"sleetAccumulationLwe"`
	SleetIntensity           int         `json:"sleetIntensity"`
	SnowAccumulation         int         `json:"snowAccumulation"`
	SnowAccumulationLwe      int         `json:"snowAccumulationLwe"`
	SnowDepth                int         `json:"snowDepth"`
	SnowIntensity            int         `json:"snowIntensity"`
	Temperature              float64     `json:"temperature"`
	TemperatureApparent      float64     `json:"temperatureApparent"`
	UvHealthConcern          int         `json:"uvHealthConcern"`
	UvIndex                  int         `json:"uvIndex"`
	Visibility               int         `json:"visibility"`
	WeatherCode              int         `json:"weatherCode"`
	WindDirection            float64     `json:"windDirection"`
	WindGust                 float64     `json:"windGust"`
	WindSpeed                float64     `json:"windSpeed"`
}

type DarkSky_Forecast struct {
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	Hourly    DarkSky_Report `json:"hourly"`
}

type DarkSky_Report struct {
	Summary string               `json:"summary"`
	Icon    string               `json:"icon"`
	Data    []DarkSky_ReportData `json:"data"`
}

type DarkSky_ReportData struct {
	Time                time.Time `json:"time"`
	Temperature         float64   `json:"temperature"`
	WindSpeed           float64   `json:"windSpeed"`
	WindBearing         float64   `json:"windBearing"`
	WindGust            float64   `json:"windGust"`
	ApparentTemperature float64   `json:"apparentTemperature"`
	Humidity            float64   `json:"humidity"`
	UvIndex             float64   `json:"uvIndex"`
	CloudCover          float64   `json:"cloudCover"`
	PrecipIntensity     float64   `json:"precipIntensity"`
	Pressure            float64   `json:"pressure"`
	Icon                string    `json:"icon"`
}

func (s *DarkSky_ReportData) UnmarshalJSON(data []byte) error {
	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		return err
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "temperature":
			s.Temperature = v.(float64)
		case "windSpeed":
			s.WindSpeed = v.(float64)
		case "windBearing":
			s.WindBearing = v.(float64)
		case "windGust":
			s.WindGust = v.(float64)
		case "apparentTemperature":
			s.ApparentTemperature = v.(float64)
		case "humidity":
			s.Humidity = v.(float64)
		case "uvIndex":
			s.UvIndex = v.(float64)
		case "cloudCover":
			s.CloudCover = v.(float64)
		case "precipIntensity":
			s.PrecipIntensity = v.(float64)
		case "pressure":
			s.Pressure = v.(float64)
		case "icon":
			s.Icon = v.(string)
		case "time":
			s.Time = time.Unix(int64(v.(float64)), 0)
		}
	}

	return nil
}
