package main

import (
	"flag"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var service Service

func main() {

	var (
		httpAddr  = flag.String("http.addr", ":6066", "Address for HTTP server")
		verbosity = flag.String("verbosity", "info", "Log verbosity")
		jsonFile  = flag.String("file", "", "Local json file to use instead of weather api")
	)

	flag.Parse()
	{
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
		lvl, err := log.ParseLevel(*verbosity)
		if err == nil {
			log.SetLevel(lvl)
		}
	}
	apiKey := os.Getenv("DARKSKY_APIKEY")
	if apiKey == "" {
		log.Fatal("missing required environment variable DARKSKY_APIKEY")
		os.Exit(1)
	}
	service = Service{apiKey, *jsonFile}

	r := mux.NewRouter()
	r.HandleFunc("/forecast/", WeatherHandler)

	loggedRouter := LoggingMiddlewar(r)

	if err := http.ListenAndServe(*httpAddr, loggedRouter); err != nil {
		log.WithField("status", "fatal").WithError(err).Fatal("fatal error")
		os.Exit(1)
	}
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	asl := r.URL.Query().Get("asl")
	coord := r.URL.Query().Get("coord") // 50.993290,4.869030 // reverse coordinates ...
	format := r.URL.Query().Get("format")

	result, err := service.GetForecast(coord, asl)
	if err != nil {
		log.WithFields(log.Fields{"result": result}).WithError(err).Error("error while getting forecast")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Connection", "close")
		w.Header().Set("Transfer-Encoding", "chunked")
		if format == "json" {
			w.Header().Set("Content-Type", "text/json")
			service.WriteJSON(w, result)
		} else {
			w.Header().Set("Content-Type", "text/xml")
			service.WriteXML(w, result)
		}
	}
}
