# loxone Weather proxy 

Written in Go.

This proxies DarkSky API data to Loxone weather service in a format it can understead.
Internally loxone uses meteoblue API for it's weather services and the response format is the same.

**! Use this for educational purposes only !**

## Requirements

- DarkSky API key
- Internal DNS server which Loxone uses. Forward `weather.loxone.com` to this server.

## Install

```.env
docker run -p 6066:6066 -e DARKSKY_APIKEY=XXXX samvdb/loxone-weather-proxy
```

