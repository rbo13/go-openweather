package openweather_test

import (
	"os"
	"testing"

	openweather "github.com/whaangbuu/go-openweather"
)

func TestGetByCityName(t *testing.T) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
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
