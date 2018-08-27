package openweather_test

import (
	"os"
	"testing"

	openweather "github.com/whaangbuu/go-openweather"
)

var apiKey string = os.Getenv("OPENWEATHER_API_KEY")

func TestGetByCityName(t *testing.T) {
	client := openweather.NewClient(apiKey)

	weatherData, err := client.GetWeatherByCityName("Cebu City")

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil {
		t.Error("Weather Data is nil")
	}

	t.Log(weatherData)
}

func TestGetByCityID(t *testing.T) {
	client := openweather.NewClient(apiKey)
	weatherData, err := client.GetWeatherByCityID(2172797)

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil {
		t.Error("Weather Data is nil")
	}

	t.Log(weatherData)
}

func TestGetByCoordinates(t *testing.T) {
	client := openweather.NewClient(apiKey)
	weatherData, err := client.GetWeatherByCoordinates(openweather.Coordinates{Latitude: 10.3157, Longitude: 123.885})

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil {
		t.Error("Weather Data is nil")
	}

	t.Log(weatherData)
}

func TestGetByZipCode(t *testing.T) {
	client := openweather.NewClient(apiKey)

	weatherData, err := client.GetWeatherByZipCode("6000", "PH")

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil {
		t.Error("Weather Data is nil")
	}

	t.Log(weatherData)
}

func TestGetForecastByCityName(t *testing.T) {
	client := openweather.NewClient(apiKey)

	forecastData, err := client.GetForecastByCityName("Cebu City")

	if err != nil {
		t.Error(err)
	}

	if forecastData == nil {
		t.Error("Forecast Data is nil")
	}

	t.Log(forecastData)
}
