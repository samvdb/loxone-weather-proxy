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

	icon := Clear

	switch code {
	case 0:
		icon = 0
	case 1000:
		icon = Clear
	case 1100:
		icon = Bright
	case 1101:
		icon = Cloudy
	case 1102:
		icon = VeryCloudy
	case 1001:
		icon = Overcast
	case 2000:
		icon = Fog
	case 2100:
		icon = LowFog
	case 4000:
		icon = Drizzle
	case 4001:
		icon = Rain
	case 4200:
		icon = LightRain
	case 4201:
		icon = HeavyRain
	case 5000:
		icon = HeavySnow
	case 5001:
		icon = LightSnow
	case 5100:
		icon = Snow
	case 5101:
		icon = StrongSnowShowers
	case 6000:
		icon = LightFreezingRain
	case 6001:
		icon = HeavyFreezingRain
	case 6200:
		icon = LightFreezingRain
	case 6201:
		icon = HeavyFreezingRain
	case 7000:
		icon = Sleet
	case 7101:
		icon = HeavySleet
	case 7102:
		icon = LightSleet
	case 8000:
		icon = Thunderstorm
	default:
		icon = 0
	}

	return icon
}
