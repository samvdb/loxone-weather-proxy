package main

// https://docs.tomorrow.io/reference/data-layers-weather-codes
//"0": "Unknown",
//	"1000": "Clear, Sunny",
//	"1100": "Mostly Clear",
//	"1101": "Partly Cloudy",
//	"1102": "Mostly Cloudy",
//	"1001": "Cloudy",
//	"2000": "Fog",
//	"2100": "Light Fog",
//	"4000": "Drizzle",
//	"4001": "Rain",
//	"4200": "Light Rain",
//	"4201": "Heavy Rain",
//	"5000": "Snow",
//	"5001": "Flurries",
//	"5100": "Light Snow",
//	"5101": "Heavy Snow",
//	"6000": "Freezing Drizzle",
//	"6001": "Freezing Rain",
//	"6200": "Light Freezing Rain",
//	"6201": "Heavy Freezing Rain",
//	"7000": "Ice Pellets",
//	"7101": "Heavy Ice Pellets",
//	"7102": "Light Ice Pellets",
//	"8000": "Thunderstorm"

func getWeatherCondition(code int) int {
	switch code {
	case 0:
		return 0
	case 1000:
		return ClearSky
	case 1100:
		return ClearSky
	case 1101:
		return PartlyCloudy
	case 1102:
		return MostlyCloudy
	case 1001:
		return MostlyCloudy
	case 2000:
		return FogLowStratusClouds
	case 2100:
		return FogLowStratusClouds
	case 4000:
		return OvercastWithRain
	case 4001:
		return OvercastWithRain
	case 4200:
		return LightRainThunderstormsLikely
	case 4201:
		return HeavyRainThunderstormsLikely
	case 5000:
		return MostlyCloudyAndCirrus
	case 5001:
		return OvercastWithRain
	case 5100:
		return OvercastWithSnow
	case 5101:
		return StormWithHeavySnow
	case 6000:
		return OvercastWithMixtureOfSnowAndRain
	case 6001:
		return OvercastWithMixtureOfSnowAndRain
	case 6200:
		return OvercastWithMixtureOfSnowAndRain
	case 6201:
		return MixedWithSnowShowers
	case 7000:
		return ClearButHazyWithCirrus
	case 7101:
		return ClearButHazyWithCirrus
	case 7102:
		return ClearButHazyWithCirrus
	case 8000:
		return HeavyRainThunderstormsLikely
	default:
		return 0
	}
}
