### GO-OPENWEATHER


`go-openweather` is a library/wrapper for the [OpenWeather API](https://openweathermap.org/) that is written in [Go](https://golang.org/).


# Features:

Has the basic access to the OpenWeather API:

  - Get the weather by City name, City ID, Coordinates, Zip Code
  - Get the Forecast ( Work-in Progress ) 


# Installation:
```sh
$ go get -u github.com/whaangbuu/go-openweather
```

# How to use:
```sh
client := openweather.NewClient("OPENWEATHER_API_KEY")

// Get Current Weather By City name
weatherData, err := client.GetWeatherByCityName("London")

// validation removed for brevity
log.Print(weatherData)
```

# TODO:
 - [ ] Get Forecast Data
 - [ ] Add more test

 # LICENSE:
 MIT