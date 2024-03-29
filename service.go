package main

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/httpcache"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var licenseExpiryDate = time.Now().AddDate(100, 0, 0)

type Service struct {
	apiKey    string
	jsonFile  string
	cacheTime int
}

func (s *Service) loadFromJson(file string) (interface{}, error) {
	log.Debug().Str("file", file).Msg("using json file")
	if _, err := os.Stat(file); err == nil {
		jsonData, err := os.Open(file)
		if err != nil {
			log.Err(err).Msg("cannot load file")
			return nil, err
		}

		defer jsonData.Close()
		var result DarkSky_Forecast
		byteData, _ := ioutil.ReadAll(jsonData)
		err = json.Unmarshal(byteData, &result)
		if err != nil {
			log.Err(err).Msg("cannot unmarshall json file")
			return nil, err
		}
		return &result, nil
	}
	return nil, fmt.Errorf("file %s does not exist", file)
}

func (s *Service) GetForecast(coord string, asl string) (interface{}, error) {

	log.Info().Str("asl", asl).
		Str("coord", coord).
		Msg("received request")

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
	log.Info().Float64("longitude", longitude).Float64("latitude", latitude).Msg("request report from tomorrow API")
	result, err := s.downloadReport(longitude, latitude)

	return result, err
}

func (s *Service) downloadReport(longitude, latitude float64) (*Tomorrow_Forecast, error) {
	request, err := http.NewRequest("GET", "https://api.tomorrow.io/v4/weather/forecast", nil)
	if err != nil {
		return nil, err
	}
	q := request.URL.Query()
	q.Add("location", fmt.Sprintf("%.3f,%.3f", latitude, longitude))
	q.Add("apikey", s.apiKey)
	q.Add("timesteps", "hourly")
	q.Add("units", "metric")
	decoded, _ := url.QueryUnescape(q.Encode())
	request.URL.RawQuery = decoded

	client := &http.Client{}
	ttl := time.Minute * time.Duration(s.cacheTime)
	_, err = httpcache.NewWithInmemoryCache(client, true, ttl)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	//dump, _ := httputil.DumpResponse(resp, true)
	//log.Debug().Msg("api response received")
	//fmt.Printf("%s", dump)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid response with code %d, error: %s", resp.StatusCode, resp.Status)
	}
	var result Tomorrow_Forecast
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&result)

	return &result, nil
}

func (s *Service) WriteJSON(w http.ResponseWriter, result interface{}) {
	report, ok := result.(*Tomorrow_Forecast)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("cannot cast %v to Tomorrow_Forecast", result)
		return
	}

	response, err := json.Marshal(report)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Interface("result", report).Err(err).Msg("cannot marshall object")
		return
	}

	log.Warn().Msg("JSON Format is for debugging only, it does not have converted values!")

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Service) WriteXML(w http.ResponseWriter, result interface{}) {
	report, ok := result.(*Tomorrow_Forecast)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("cannot cast %v to Tomorrow_Forecast", result)
		return
	}

	xml := "<?xml version=\"1.0\"?>"
	xml += fmt.Sprintf("<metdata_feature_collection p=\"m\" valid_untill=\"%d-%d-%d\">", licenseExpiryDate.Year(), licenseExpiryDate.Month(), licenseExpiryDate.Day())

	for _, hourly := range report.Timelines.Hourly {
		xml += "<metdata>"
		xml += fmt.Sprintf("<timepoint>%s</timepoint>", hourly.Time.Format("2006-01-02T15:04:05"))
		xml += fmt.Sprintf("<TT>%.1f</TT>", hourly.Values.Temperature)
		xml += fmt.Sprintf("<FF>%.1f</FF>", hourly.Values.WindSpeed*1000/3600)
		windBearing := hourly.Values.WindDirection - 180
		if windBearing < 0 {
			windBearing += 360
		}
		xml += fmt.Sprintf("<DD>%.0f</DD>", windBearing)
		xml += fmt.Sprintf("<RR1H>%0.1f</RR1H>", hourly.Values.RainIntensity)
		xml += fmt.Sprintf("<PP0>%.0f</PP0>", hourly.Values.PressureSurfaceLevel)
		xml += fmt.Sprintf("<RH>%.0f</RH>", hourly.Values.Humidity*100)
		xml += fmt.Sprintf("<HI>%.1f</HI>", hourly.Values.TemperatureApparent)
		xml += fmt.Sprintf("<RAD>%4.0f</RAD>", hourly.Values.UVIndex*100)
		xml += fmt.Sprintf("<WW>%d</WW>", s.fixIcon(hourly))
		xml += fmt.Sprintf("<FFX>%.1f</FFX>", hourly.Values.WindGust*100/3600)
		xml += "<LC>0</LC>"
		xml += fmt.Sprintf("<MC>%.0f</MC>", hourly.Values.CloudCover*100)
		xml += "<HC>0</HC>"
		xml += fmt.Sprintf("<RAD4C>%.0f</RAD4C>", hourly.Values.UVIndex)
		xml += "</metdata>"
	}
	xml += "</metdata_feature_collection>\n"

	//log.Debug().Str("xml", xml).Msg("xml data")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(xml))
}

func (s *Service) fixIcon(report Tomorrow_Report) int {
	//return fixDarkSky2(report)
	return getWeatherCondition(report.Values.WeatherCode)
}
