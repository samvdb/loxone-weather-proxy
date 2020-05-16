package main

import (
	. "github.com/sirupsen/logrus"
	"time"
)
// A machine-readable text summary of this data point, suitable for selecting an icon for display.
//If defined, this property will have one of the following values: clear-day, clear-night, rain, snow, sleet, wind, fog, cloudy, partly-cloudy-day, or partly-cloudy-night. (Developers should ensure that a sensible default is defined, as additional values, such as hail, thunderstorm, or tornado, may be defined in the future.)


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
func fixDarkSky1(report DarkSky_ReportData) int {
	var id = 0
	switch report.Icon {
	case "clear-day", "clear-night":
		id = 1
	case "rain":
		id = 23
	case "snow":
		id = 24
	case "sleet":
		id = 35
	case "fog":
		id = 16
	case "cloudy", "partly-cloudy-day", "partly-cloudy-night":
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
	WithFields(Fields{"string_icon": report.Icon, "converted_id": id, "time": report.Time.Format(time.RFC3339)}).Debug("convert icon")
	return id
}


func fixDarkSky2(report DarkSky_ReportData) int {
	var id = 0
	switch report.Icon {
	case "clear-day", "clear-night":
		id = 1
	case "rain":
		if report.PrecipIntensity < 0.5 {
			id = 10
		} else if report.PrecipIntensity <= 4 {
			id = 11
		} else {
			id = 12
		}
	case "snow":
		id = 21
	case "sleet":
		id = 26
	case "fog":
		id = 6
	case "cloudy":
		id = 3
	case "partly-cloudy-day", "partly-cloudy-night":
		id = 2
	case "hail":
		id = 15
	case "thunderstorm":
		id = 18
	default:
		id = 1
	}

	switch id {
	case 2,3:
		if report.CloudCover < 0.125 {
			id = 2
		} else if report.CloudCover < 0.5 {
			id = 3
		} else {
			id = 4
		}

	}
	WithFields(Fields{"string_icon": report.Icon, "converted_id": id, "time": report.Time.Format(time.RFC3339)}).Debug("convert icon")
	return id
}


