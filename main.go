package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"encoding/json"
	"net/http"
	"os"
)

var service Service

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {

	var (
		httpAddr = flag.String("http.addr", ":6066", "Address for HTTP server")
		debug    = flag.Bool("debug", false, "sets log level to debug")
		jsonFile  = flag.String("file", "", "Local json file to use instead of weather api")
	)

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Info().Str("version", version).Str("commit", commit).Str("date", date).Msg("")

	apiKey := os.Getenv("DARKSKY_APIKEY")
	if apiKey == "" {
		log.Fatal().Msg("missing required environment variable DARKSKY_APIKEY")
		os.Exit(1)
	}
	service = Service{apiKey, *jsonFile}

	r := mux.NewRouter()
	c := LogHttp(log.Logger)
	r.Handle("/forecast/", c.Then(http.HandlerFunc(WeatherHandler)))
	
	r.Handle("/interface/api", c.Then(http.HandlerFunc(ApiHandler)))
	r.Handle("/interface/api.php", c.Then(http.HandlerFunc(ApiHandler)))
	if err := http.ListenAndServe(*httpAddr, r); err != nil {
		log.Fatal().Str("status", "fatal").Err(err).Msg("fatal error")
		os.Exit(1)
	}
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	asl := r.URL.Query().Get("asl")
	coord := r.URL.Query().Get("coord") // 50.993290,4.869030 // reverse coordinates ...
	format := r.URL.Query().Get("format")

	result, err := service.GetForecast(coord, asl)
	if err != nil {
		log.Error().Interface("result", result).Err(err).Msg("error while getting forecast")
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


func ApiHandler(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("received /interface/api call")
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Connection", "close")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "text/json")
		result := map[string]string{
			"1": "2100-01-01",
			"2": "2100-01-01",
		}
		response, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error().Interface("result", result).Err(err).Msg("cannot marshall object")
			return
		}
	
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	
}
