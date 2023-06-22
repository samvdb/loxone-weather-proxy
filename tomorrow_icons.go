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

	icon := ClearSky
	switch code {
	case 0:
		icon = 0
	case 1000:
		icon = ClearSky
	case 1100:
		icon = ClearFewCirrus
	case 1101:
		icon = PartlyCloudy
	case 1102:
		icon = MostlyCloudy
	case 1001:
		icon = Overcast
	case 2000:
		icon = FogLowStratusClouds
	case 2100:
		icon = FogLowStratusCloudsWithFewCirrus
	case 4000:
		icon = OvercastWithRain
	case 4001:
		icon = OvercastWithRain
	case 4200:
		icon = OvercastWithLightRain
	case 4201:
		icon = OvercastWithHeavyRain
	case 5000:
		icon = OvercastWithSnow
	case 5001:
		icon = OvercastWithMixtureOfSnowAndRain
	case 5100:
		icon = OvercastWithLightSnow
	case 5101:
		icon = OvercastWithHeavySnow
	// checked above
	case 6000:
		icon = OvercastWithMixtureOfSnowAndRain
	case 6001:
		icon = OvercastWithMixtureOfSnowAndRain
	case 6200:
		icon = OvercastWithMixtureOfSnowAndRain
	case 6201:
		icon = MixedWithSnowShowers
	case 7000:
		icon = OvercastWithMixtureOfSnowAndRain
	case 7101:
		icon = OvercastWithMixtureOfSnowAndRain
	case 7102:
		icon = OvercastWithMixtureOfSnowAndRain
	case 8000:
		icon = HeavyRainThunderstormsLikely
	default:
		icon = 0
	}

	return icon
}
