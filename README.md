

[![Build Status](https://travis-ci.com/samvdb/loxone-weather-proxy.svg?branch=master)](https://travis-ci.com/samvdb/loxone-weather-proxy)

# Overview 

Written in Go.

This proxies Tomorrow API data to Loxone weather service in a format it can understand.
Internally loxone uses meteoblue API for it's weather services and the response format is the same.

**! Use this for educational purposes only !**

## Requirements

- Tomorrow API key
- Internal DNS server which Loxone uses. Forward `weather.loxone.com` to this server.

## Install

```.env
docker run -p 6066:6066 -e TOMORROW_APIKEY=XXXX ghcr.io/samvdb/loxone-weather-proxy
```

## Weather type

### Weather type ##

Return value  | Desc                 |
------------- | -------------------  |
1             | No clouds            |
2             | Clear                |
3             | Scattered clouds     |
4             | Heavy cloud cover    |
5             | Overcast             |
6             | Fog                  |
7             | Havy fog             |
8             | N/A                  |
9             | N/A                  |
10            | Light Rain           |
11            | Rain                 |
12            | Heavy Rain           |
13            | Drizzle              |
14            | Sleet                |
15            | Heavy Freezing Rain  |
16            | Light Shower         |
17            | Heavy rain showers   |
18            | Thunderstorms        |
19            | Strong thunderstorms |
20            | Light snow           |
21            | Snow                 |
22            | Heavy snow           |
23            | Light snow showers   |
24            | heavy snow showers   |
25            | Light sleet          |
26            | sleet                |
27            | heavy sleet          |
28            | Light sleet showers  |
29            | heavy sleet showers  |
