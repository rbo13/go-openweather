package openweather_test

import (
	"os"
	"testing"

	openweather "github.com/rbo13/go-openweather"
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

func TestGetForecastByCityID(t *testing.T) {
	client := openweather.NewClient(apiKey)

	forecastData, err := client.GetForecastByCityID("524901")

	if err != nil {
		t.Error(err)
	}

	if forecastData == nil {
		t.Error("Forecast Data is nil")
	}

	t.Log(forecastData)
}

func TestGetForecastByCoordinates(t *testing.T) {
	client := openweather.NewClient(apiKey)

	forecastData, err := client.GetForecastByCoordinates(openweather.Coordinates{Latitude: 35, Longitude: 139})

	if err != nil {
		t.Error(err)
	}

	if forecastData == nil {
		t.Error("Forecast Data is nil")
	}

	t.Log(forecastData)
}

func TestGetForecastByZipCode(t *testing.T) {
	client := openweather.NewClient(apiKey)

	forecastData, err := client.GetForecastByZipCode("6000", "PH")

	if err != nil {
		t.Error(err)
	}

	if forecastData == nil {
		t.Error("Forecast Data is nil")
	}

	t.Log(forecastData)
}

func TestGetDailyForecastByCityName(t *testing.T) {
	client := openweather.NewClient(apiKey)

	dailyForecastData, err := client.GetDailyForecastByCityName("Cebu City", "", "")

	if err != nil {
		t.Error(err)
	}

	if dailyForecastData == nil {
		t.Error("Daily Forecast Data is nil")
	}

	t.Log(dailyForecastData)
}

func TestGetDailyForecastByCityID(t *testing.T) {
	client := openweather.NewClient(apiKey)

	dailyForecastData, err := client.GetDailyForecastByCityID("524901", "7")

	if err != nil {
		t.Error(err)
	}

	if dailyForecastData == nil {
		t.Error("Daily Forecast Data is nil")
	}

	t.Log(dailyForecastData)
}

func TestGetDailyForecastByCoordinates(t *testing.T) {
	client := openweather.NewClient(apiKey)

	dailyForecastData, err := client.GetDailyForecastByCoordinates(openweather.Coordinates{Latitude: 35, Longitude: 139}, "")

	if err != nil {
		t.Error(err)
	}

	if dailyForecastData == nil {
		t.Error("Daily Forecast Data is nil")
	}

	t.Log(dailyForecastData)
}
