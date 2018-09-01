package openweather_test

import (
	"os"
	"testing"

	openweather "github.com/rbo13/go-openweather"
	"github.com/rbo13/go-openweather/coords"
	"github.com/rbo13/go-openweather/forecast"

	"github.com/rbo13/go-openweather/weather"
)

var apiKey string = os.Getenv("OPENWEATHER_API_KEY")

const baseURL string = "https://api.openweathermap.org/data/2.5"

var w *weather.Weather
var f *forecast.Forecast

func init() {
	w = weather.NewWeather(baseURL)
	f = forecast.NewForecast(baseURL)
}

func TestGetByCityName(t *testing.T) {
	openweather := openweather.New(w, f)

	weatherData, err := openweather.Weatherer.GetByCityName("Cebu City")
	if err != nil {
		t.Error(err)
	}

	if weatherData == nil {
		t.Error("Weather is nil")
	}

	forecastData, err := openweather.Forecaster.GetByCityName("Cebu City")

	if err != nil {
		t.Error(err)
	}

	if forecastData == nil {
		t.Error("Forecast is nil")
	}

	t.Log(weatherData)
	t.Log(forecastData)
}

func TestGetByCityID(t *testing.T) {
	weatherData, err := w.GetByCityID(2172797)
	forecastData, err := f.GetByCityID(2172797)

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
	weatherData, err := w.GetByCoordinates(coords.Coordinates{Latitude: 10.3157, Longitude: 123.885})
	forecastData, err := f.GetByCoordinates(coords.Coordinates{Latitude: 10.3157, Longitude: 123.885})

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
	weatherData, err := w.GetByZipCode("6000", "PH")
	forecastData, err := f.GetByZipCode("6000", "PH")

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil || forecastData == nil {
		t.Error("ERROR")
	}

	t.Log(weatherData)
	t.Log(forecastData)
}
