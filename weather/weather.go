package weather

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/rbo13/go-openweather/client"
	"github.com/rbo13/go-openweather/coords"
)

// Weather ...
type Weather struct {
	client.Client
	Url         string
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

func NewWeather(baseUrl string) *Weather {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &Weather{
		Client: client.Client{
			APIKey: os.Getenv("OPENWEATHER_API_KEY"),
			HTTPClient: &http.Client{
				Timeout:   10 * time.Second,
				Transport: netTransport,
			},
		},
		Url:         baseUrl,
		WeatherData: new(WeatherData),
	}
}

func (w *Weather) GetByCityName(cityName string) (*WeatherData, error) {
	apiURL := fmt.Sprintf(w.Url+"/weather?q=%s&appid=%s", cityName, w.Client.APIKey)
	err := w.Client.Request("GET", apiURL, &w.WeatherData)

	if err != nil {
		return nil, err
	}
	return w.WeatherData, nil
}

func (w *Weather) GetByCityID(cityID int64) (*WeatherData, error) {
	apiURL := fmt.Sprintf(w.Url+"/weather?id=%d&appid=%s", cityID, w.Client.APIKey)
	err := w.Client.Request("GET", apiURL, &w.WeatherData)

	if err != nil {
		return nil, err
	}
	return w.WeatherData, nil
}

func (w *Weather) GetByCoordinates(coords coords.Coordinates) (*WeatherData, error) {
	apiURL := fmt.Sprintf(w.Url+"/weather?lat=%g&lon=%g&appid=%s", coords.Latitude, coords.Longitude, w.Client.APIKey)
	err := w.Client.Request("GET", apiURL, &w.WeatherData)

	if err != nil {
		return nil, err
	}

	return w.WeatherData, nil
}

func (w *Weather) GetByZipCode(zipCode, countryCode string) (*WeatherData, error) {
	apiURL := fmt.Sprintf(w.Url+"/weather?zip=%s,%s&appid=%s", zipCode, countryCode, w.Client.APIKey)
	err := w.Client.Request("GET", apiURL, &w.WeatherData)

	if err != nil {
		return nil, err
	}
	return w.WeatherData, nil
}
