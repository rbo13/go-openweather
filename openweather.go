package openweather

import (
	"github.com/rbo13/go-openweather/coords"
)

// WeatherForecaster ...
type WeatherForecaster interface {
	GetByCityName(cityName string) (*Openweather, error)
	GetByCityID(cityID int64) (*Openweather, error)
	GetByCoordinates(coords coords.Coordinates) (*Openweather, error)
	GetByZipCode(zipCode, countryCode string) (*Openweather, error)
}

// WeatherForecastService ...
type WeatherForecastService interface {
	WeatherForecaster
}

// Openweather ...
type Openweather struct {
	WeatherForecaster
}

// New ...
func New(wf WeatherForecaster) WeatherForecastService {
	return &Openweather{WeatherForecaster: wf}
}

// type (
// 	Weatherer interface {
// 		GetByCityName(cityName string) (*weather.WeatherData, error)
// 		GetByCityID(cityID int64) (*weather.WeatherData, error)
// 		GetByCoordinates(coords coords.Coordinates) (*weather.WeatherData, error)
// 		GetByZipCode(zipCode, countryCode string) (*weather.WeatherData, error)
// 	}

// 	Forecaster interface {
// 		GetByCityName(cityName string) (*forecast.ForecastData, error)
// 		GetByCityID(cityID int64) (*forecast.ForecastData, error)
// 		GetByCoordinates(coords coords.Coordinates) (*forecast.ForecastData, error)
// 		GetByZipCode(zipCode, countryCode string) (*forecast.ForecastData, error)
// 	}
// 	Openweather struct {
// 		Weatherer
// 		Forecaster
// 	}
// )

// func New(weather Weatherer, forecaster Forecaster) *Openweather {
// 	return &Openweather{
// 		Weatherer:  weather,
// 		Forecaster: forecaster,
// 	}
// }

// func (ow Openweather) GetWeatherByCityName(cityName string) (*WeatherData, error) {
// 	return ow.Weather.GetByCityName(cityName)
// }
