package main

// https://github.com/sarnau/Inside-The-Loxone-Miniserver/blob/master/Code/LoxoneWeather.py
//
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
const (
	ClearSky                          = 1
	ClearFewCirrus                    = 2
	ClearWithCirrus                   = 3
	ClearWithFewLowClouds             = 4
	ClearWithFewLowCloudsAndFewCirrus = 5
	ClearWithFewLowCloudsAndCirrus    = 6
	PartlyCloudy                      = 7
	PartlyCloudyAndFewCirrus          = 8
	PartlyCloudyAndCirrus             = 9
	MixedWithSomeThunderstormClouds   = 10
	MixedWithFewCirrusAndThunderstorm = 11
	MixedWithCirrusAndThunderstorm    = 12
	ClearButHazy                      = 13
	ClearButHazyWithFewCirrus         = 14
	ClearButHazyWithCirrus            = 15
	FogLowStratusClouds               = 16
	FogLowStratusCloudsWithFewCirrus  = 17
	FogLowStratusCloudsWithCirrus     = 18
	MostlyCloudy                      = 19
	MostlyCloudyAndFewCirrus          = 20
	MostlyCloudyAndCirrus             = 21
	Overcast                          = 22
	OvercastWithRain                  = 23
	OvercastWithSnow                  = 24
	OvercastWithHeavyRain             = 25
	OvercastWithHeavySnow             = 26
	RainThunderstormsLikely           = 27
	LightRainThunderstormsLikely      = 28
	StormWithHeavySnow                = 29
	HeavyRainThunderstormsLikely      = 30
	MixedWithShowers                  = 31
	MixedWithSnowShowers              = 32
	OvercastWithLightRain             = 33
	OvercastWithLightSnow             = 34
	OvercastWithMixtureOfSnowAndRain  = 35
)
