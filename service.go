package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var licenseExpiryDate = time.Now().AddDate(100, 0, 0)

type Service struct {
	apiKey   string
	jsonFile string
}

func (s *Service) loadFromJson(file string) (interface{}, error) {
	logger := log.WithField("file", file)
	if _, err := os.Stat(file); err == nil {
		logger.Info("loading weather from json file")
		jsonData, err := os.Open(file)
		if err != nil {
			return nil, err
		}

		logger.Debug("file read complete")
		defer jsonData.Close()
		var result DarkSky_Forecast
		byteData, _ := ioutil.ReadAll(jsonData)
		err = json.Unmarshal(byteData, &result)
		if err != nil {
			log.WithError(err).Error("cannot unmarshall json file")
			return nil, err
		}
		return &result, nil
	}
	return nil, fmt.Errorf("file %s does not exist", file)
}

func (s *Service) GetForecast(coord string, asl string) (interface{}, error) {

	if s.jsonFile != "" {
		return s.loadFromJson(s.jsonFile)
	}

	ce := strings.Split(coord, ",")
	if len(ce) != 2 {
		return nil, fmt.Errorf("coords are not valid: %s", coord)
	}
	// loxone reverses coordinates
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
	log.WithField("response", result).Debug("api response received")

	return &result, nil
}


func (s *Service) WriteJSON(w http.ResponseWriter, result interface{}) {
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

	log.Warn("JSON Format is for debugging only, it does not have converted values!")

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



// Generate an icon for Loxone based on the Meteoblue picto-codes
// <https://content.meteoblue.com/en/help/standards/symbols-and-pictograms>
//  1	Clear, cloudless sky (Loxone: Wolkenlos)
//  2	Clear, few cirrus (Loxone: Wolkenlos)
//  3	Clear with cirrus (Loxone: Heiter)
//  4	Clear with few low clouds (Loxone: Heiter)
//  5	Clear with few low clouds and few cirrus (Loxone: Heiter)
//  6	Clear with few low clouds and cirrus (Loxone: Heiter)
//  7	Partly cloudy (Loxone: Heiter)
//  8	Partly cloudy and few cirrus (Loxone: Heiter)
//  9	Partly cloudy and cirrus (Loxone: Wolkig)
// 10	Mixed with some thunderstorm clouds possible (Loxone: Wolkig)
// 11	Mixed with few cirrus with some thunderstorm clouds possible (Loxone: Wolkig)
// 12	Mixed with cirrus and some thunderstorm clouds possible (Loxone: Wolkig)
// 13	Clear but hazy (Loxone: Wolkenlos)
// 14	Clear but hazy with few cirrus (Loxone: Heiter)
// 15	Clear but hazy with cirrus (Loxone: Heiter)
// 16	Fog/low stratus clouds (Loxone: Nebel)
// 17	Fog/low stratus clouds with few cirrus (Loxone: Nebel)
// 18	Fog/low stratus clouds with cirrus (Loxone: Nebel)
// 19	Mostly cloudy (Loxone: Stark bewölkt)
// 20	Mostly cloudy and few cirrus (Loxone: Stark bewölkt)
// 21	Mostly cloudy and cirrus (Loxone: Stark bewölkt)
// 22	Overcast (Loxone: Bedeckt)
// 23	Overcast with rain (Loxone: Regen)
// 24	Overcast with snow (Loxone: Schneefall)
// 25	Overcast with heavy rain (Loxone: Starker Regen)
// 26	Overcast with heavy snow (Loxone: Starker Schneefall)
// 27	Rain, thunderstorms likely (Loxone: Kräftiges Gewitter)
// 28	Light rain, thunderstorms likely (Loxone: Gewitter)
// 29	Storm with heavy snow (Loxone: Starker Schneeschauer)
// 30	Heavy rain, thunderstorms likely (Loxone: Kräftiges Gewitter)
// 31	Mixed with showers (Loxone: Leichter Regenschauer)
// 32	Mixed with snow showers (Loxone: Leichter Schneeschauer)
// 33	Overcast with light rain (Loxone: Leichter Regen)
// 34	Overcast with light snow (Loxone: Leichter Schneeschauer)
// 35	Overcast with mixture of snow and rain (Loxone: Schneeregen)
func (s *Service) fixIcon(report DarkSky_ReportData) int {
	var id = 0
	switch report.Icon {
	case "clear-day","clear-night":
		id = 1
	case "rain":
		id = 23
	case "snow":
		id = 24
	case "sleet":
		id = 35
	case "fog":
		id = 16
	case "cloudy","partly-cloudy-day","partly-cloudy-night":
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
	log.WithFields(log.Fields{"string_icon": report.Icon, "converted_id": id}).Debug("convert icon")
	return id
}