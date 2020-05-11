package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var licenseExpiryDate = time.Now().AddDate(100, 0, 0)

type Service struct {
	apiKey string
}

func (s *Service) GetForecast(coord string, asl string) (interface{}, error) {
	ce := strings.Split(coord, ",")
	if len(ce) != 2 {
		return nil, fmt.Errorf("coords are not valid: %s", coord)
	}
	longitude, err := strconv.ParseFloat(ce[0], 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse longitude: %s", ce[0])
	}
	latitude, err := strconv.ParseFloat(ce[1], 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse latitude: %s", ce[1])
	}
	log.WithFields(log.Fields{
		"longitude": longitude,
		"latitude":  latitude,
	}).Info("request report from darkSky API")
	result, err := s.downloadReport(longitude, latitude)

	return result, err
}

func (s *Service) downloadReport(longitude, latitude float64) (*DarkSky_Forecast, error) {
	path := fmt.Sprintf("https://api.darksky.net/forecast/%s/%.3f,%.2f", s.apiKey, latitude, longitude)
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	q := request.URL.Query()
	q.Add("extend", "hourly")
	q.Add("lang", "en")
	q.Add("units", "ca")
	request.URL.RawQuery = q.Encode()

	resp, err := http.Get(request.URL.String())
	if err != nil {
		return nil, err
	}
	var result DarkSky_Forecast
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)

	return &result, nil
}

func (s *Service) fixIcon(report DarkSky_ReportData) int {
	var id = 0
	switch report.Icon {
	case "clear-day":
	case "clear-night":
		id = 1
	case "rain":
		id = 23
	case "snow":
		id = 24
	case "sleet":
		id = 35
	case "fog":
		id = 16
	case "cloudy":
	case "partly-cloudy-day":
	case "partly-cloudy-night":
		id = 7
	case "hail":
		id = 35
	case "thunderstorm":
		id = 28
	default:
		id = 7
	}

	switch id {
	case 7:
		if report.CloudCover < 0.125 {
			id = 1
		} else if report.CloudCover < 0.5 {
			id = 3
		} else if report.CloudCover < 0.75 {
			id = 9
		} else if report.CloudCover < 0.875 {
			id = 19
		} else {
			id = 22
		}

	case 23:
		if report.PrecipIntensity < 0.5 {
			id = 33
		} else if report.PrecipIntensity <= 4 {
			id = 23
		} else {
			id = 25
		}
	}
	return id
}

func (s *Service) WriteCSV(w http.ResponseWriter, result interface{}) {
	report, ok := result.(*DarkSky_Forecast)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("cannot cast %v to DarkSky_Forecast", result)
		return
	}

	response, err := json.Marshal(report)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithField("result", report).WithError(err).Error("cannot marshall object")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Service) WriteXML(w http.ResponseWriter, result interface{}) {
	report, ok := result.(*DarkSky_Forecast)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("cannot cast %v to DarkSky_Forecast", result)
		return
	}

	xml := "<?xml version=\"1.0\"?>"
	xml += fmt.Sprintf("<metdata_feature_collection p=\"m\" valid_untill=\"%d-%d-%d\">", licenseExpiryDate.Year(), licenseExpiryDate.Month(), licenseExpiryDate.Day())

	for _, hourly := range report.Hourly.Data {
		xml += "<metdata>"
		xml += fmt.Sprintf("<timepoint>%s</timepoint>", hourly.Time.Format("2006-01-02T15:04:05"))
		xml += fmt.Sprintf("<TT>%.1f</TT>", hourly.Temperature)
		xml += fmt.Sprintf("<FF>%.1f</FF>", hourly.WindSpeed*1000/3600)
		windBearing := hourly.WindBearing - 180
		if windBearing < 0 {
			windBearing += 360
		}
		xml += fmt.Sprintf("<DD>%.0f</DD>", windBearing)
		xml += fmt.Sprintf("<RR1H>%5.1f</RR1H>", hourly.PrecipIntensity)
		xml += fmt.Sprintf("<PP0>%.0f</PP0>", hourly.Pressure)
		xml += fmt.Sprintf("<RH>%.0f</RH>", hourly.Humidity*100)
		xml += fmt.Sprintf("<HI>%.1f</HI>", hourly.ApparentTemperature)
		xml += fmt.Sprintf("<RAD>%4.0f</RAD>", hourly.UvIndex*100)
		xml += fmt.Sprintf("<WW>%d</WW>", s.fixIcon(hourly))
		xml += fmt.Sprintf("<FFX>%.1f</FFX>", hourly.WindGust*100/3600)
		xml += "<LC>0</LC>"
		xml += fmt.Sprintf("<MC>%.0f</MC>", hourly.CloudCover*100)
		xml += "<HC>0</HC>"
		xml += fmt.Sprintf("<RAD4C>%.0f</RAD4C>", hourly.UvIndex)
		xml += "</metdata>"
	}
	xml += "</metdata_feature_collection>\n"

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(xml))
}
