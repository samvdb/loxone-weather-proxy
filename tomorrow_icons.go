package main

//	1	Clear, cloudless sky (Loxone: Wolkenlos)
//	2	Clear, few cirrus (Loxone: Wolkenlos)
//	3	Clear with cirrus (Loxone: Heiter)
//	4	Clear with few low clouds (Loxone: Heiter)
//	5	Clear with few low clouds and few cirrus (Loxone: Heiter)
//	6	Clear with few low clouds and cirrus (Loxone: Heiter)
//	7	Partly cloudy (Loxone: Heiter)
//	8	Partly cloudy and few cirrus (Loxone: Heiter)
//	9	Partly cloudy and cirrus (Loxone: Wolkig)
//
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
func fixTomorrow(report Tomorrow_Report) int {

	// https://docs.tomorrow.io/reference/data-layers-weather-codes
	switch report.Values.WeatherCode {
	case 1000:
		return
	}
	var id = 0
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
	return id

}
