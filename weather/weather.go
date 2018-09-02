package weather

import (
	"fmt"
	"net/http"
	"os"
	"time"

	openweather "github.com/rbo13/go-openweather"
	"github.com/rbo13/go-openweather/client"
	"github.com/rbo13/go-openweather/coords"
)

// Weather ...
type Weather struct {
	client.Client
	URL         string
	WeatherData *WeatherData
}

// WeatherData represents
// the weather information
// given by a certain parameters
type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		Pressure  float64 `json:"pressure"`
		Humidity  int     `json:"humidity"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		SeaLevel  float64 `json:"sea_level"`
		GrndLevel float64 `json:"grnd_level"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

// NewWeather ...
func NewWeather(baseURL string, transport *http.Transport) *Weather {
	return &Weather{
		Client: client.Client{
			APIKey: os.Getenv("OPENWEATHER_API_KEY"),
			HTTPClient: &http.Client{
				Timeout:   10 * time.Second,
				Transport: transport,
			},
		},
		URL:         baseURL,
		WeatherData: new(WeatherData),
	}
}

// GetByCityName ...
func (w *Weather) GetByCityName(cityName string) (*openweather.Openweather, error) {
	apiURL := fmt.Sprintf(w.URL+"/weather?q=%s&appid=%s", cityName, w.Client.APIKey)
	err := w.Client.Request("GET", apiURL, &w.WeatherData)
	if err != nil {
		return nil, err
	}
	return &openweather.Openweather{
		WeatherForecaster: w,
	}, nil
}

// GetByCityID ...
func (w *Weather) GetByCityID(cityID int64) (*openweather.Openweather, error) {
	apiURL := fmt.Sprintf(w.URL+"/weather?id=%d&appid=%s", cityID, w.Client.APIKey)
	err := w.Client.Request("GET", apiURL, &w.WeatherData)
	if err != nil {
		return nil, err
	}

	return &openweather.Openweather{
		WeatherForecaster: w,
	}, nil
}

// GetByCoordinates ...
func (w *Weather) GetByCoordinates(coords coords.Coordinates) (*openweather.Openweather, error) {
	apiURL := fmt.Sprintf(w.URL+"/weather?lat=%g&lon=%g&appid=%s", coords.Latitude, coords.Longitude, w.Client.APIKey)
	err := w.Client.Request("GET", apiURL, &w.WeatherData)
	if err != nil {
		return nil, err
	}

	return &openweather.Openweather{
		WeatherForecaster: w,
	}, nil
}

// GetByZipCode ...
func (w *Weather) GetByZipCode(zipCode, countryCode string) (*openweather.Openweather, error) {
	apiURL := fmt.Sprintf(w.URL+"/weather?zip=%s,%s&appid=%s", zipCode, countryCode, w.Client.APIKey)
	err := w.Client.Request("GET", apiURL, &w.WeatherData)
	if err != nil {
		return nil, err
	}
	return &openweather.Openweather{
		WeatherForecaster: w,
	}, nil
}
