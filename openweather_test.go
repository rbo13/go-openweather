package openweather_test

import (
	"os"
	"testing"

	openweather "github.com/rbo13/go-openweather"
)

var apiKey string = os.Getenv("OPENWEATHER_API_KEY")

const baseURL string = "https://api.openweathermap.org/data/2.5"

var weather *openweather.Weather
var forecast *openweather.Forecast

func init() {
	weather = openweather.NewWeather(baseURL)
	forecast = openweather.NewForecast(baseURL)
}

func TestGetByCityName(t *testing.T) {
	weatherData, err := weather.GetByCityName("Cebu City")
	forecastData, err := forecast.GetByCityName("Cebu City")

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil || forecastData == nil {
		t.Error("ERROR")
	}

	t.Log(weatherData)
	t.Log(forecastData)
}

func TestGetByCityID(t *testing.T) {
	weatherData, err := weather.GetByCityID(2172797)
	forecastData, err := forecast.GetByCityID(2172797)

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil || forecastData == nil {
		t.Error("Weather Data is nil")
	}

	t.Log(weatherData)
	t.Log(forecastData)
}

func TestGetByCoordinates(t *testing.T) {
	weatherData, err := weather.GetByCoordinates(openweather.Coordinates{Latitude: 10.3157, Longitude: 123.885})
	forecastData, err := forecast.GetByCoordinates(openweather.Coordinates{Latitude: 10.3157, Longitude: 123.885})

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil || forecastData == nil {
		t.Error("ERROR")
	}

	t.Log(weatherData)
	t.Log(forecastData)
}

func TestByZipCode(t *testing.T) {
	weatherData, err := weather.GetByZipCode("6000", "PH")
	forecastData, err := forecast.GetByZipCode("6000", "PH")

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil || forecastData == nil {
		t.Error("ERROR")
	}

	t.Log(weatherData)
	t.Log(forecastData)
}
