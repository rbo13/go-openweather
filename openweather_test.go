package openweather_test

import (
	"os"
	"testing"

	openweather "github.com/whaangbuu/go-openweather"
)

var apiKey string = os.Getenv("OPENWEATHER_API_KEY")

func TestGetByCityName(t *testing.T) {
	client := openweather.NewClient(apiKey)

	weatherData, err := client.GetWeatherByCityName("London")

	if err != nil {
		t.Error(err)
	}

	if weatherData == nil {
		t.Errorf("Weather Data is nil")
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
		t.Errorf("Weather Data i nil")
	}

	t.Log(weatherData)
}
