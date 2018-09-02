package openweather_test

import (
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/rbo13/go-openweather/coords"
	"github.com/rbo13/go-openweather/forecast"

	"github.com/rbo13/go-openweather/weather"
)

var apiKey string = os.Getenv("OPENWEATHER_API_KEY")

const baseURL string = "https://api.openweathermap.org/data/2.5"

var w *weather.Weather
var f *forecast.Forecast

var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 5 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 5 * time.Second,
}

func init() {
	w = weather.NewWeather(baseURL, netTransport)
	f = forecast.NewForecast(baseURL, netTransport)
}

func TestGetByCityName(t *testing.T) {
	openweatherData, err := w.GetByCityName("Cebu City")
	forecast, err := f.GetByCityName("Cebu City")

	if err != nil {
		t.Error(err)
	}

	if openweatherData == nil || forecast == nil {
		t.Error("Must not be nil")
	}

	t.Log(w.WeatherData)
	t.Log(f.ForecastData)
}

func TestGetByCityID(t *testing.T) {
	openweatherData, err := w.GetByCityID(2172797)
	forecast, err := f.GetByCityID(2172797)

	if err != nil {
		t.Error(err)
	}

	if openweatherData == nil || forecast == nil {
		t.Error("Must not be nil")
	}

	t.Log(w.WeatherData)
	t.Log(f.ForecastData)
}

func TestGetByCoordinates(t *testing.T) {
	openweatherData, err := w.GetByCoordinates(coords.Coordinates{Latitude: 10.3157, Longitude: 123.885})
	forecast, err := f.GetByCoordinates(coords.Coordinates{Latitude: 10.3157, Longitude: 123.885})

	if err != nil {
		t.Error(err)
	}

	if openweatherData == nil || forecast == nil {
		t.Error("Must not be nil")
	}

	t.Log(w.WeatherData)
	t.Log(f.ForecastData)
}

func TestGetByZipCode(t *testing.T) {
	openweatherData, err := w.GetByZipCode("6000", "PH")
	forecast, err := f.GetByZipCode("6000", "PH")

	if err != nil {
		t.Error(err)
	}

	if openweatherData == nil || forecast == nil {
		t.Error("Must not be nil")
	}

	t.Log(w.WeatherData)
	t.Log(f.ForecastData)
}
