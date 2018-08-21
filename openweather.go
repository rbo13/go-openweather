package openweather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

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
		Deg   int     `json:"deg"`
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

// Client represents
// the client request to
// openweather api
type Client struct {
	apiKey string
}

const baseURL string = "https://api.openweathermap.org/data/2.5/weather"

// https://samples.openweathermap.org/data/2.5/weather?q=London,uk&appid=b6907d289e10d714a6e88b30761fae22

// NewClient returns the Client struct
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

// GetWeatherByCityName returns the
// weather in a given city
func (c *Client) GetWeatherByCityName(cityName string) (*WeatherData, error) {
	apiURL := fmt.Sprintf(baseURL+"?q=%s&appid=%s", cityName, c.apiKey)
	var weatherData WeatherData

	response, err := requestQuery().Get(apiURL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	r := bytes.NewBuffer(b)
	err = json.NewDecoder(r).Decode(&weatherData)

	log.Print(apiURL)

	return &weatherData, nil
}

func requestQuery() *http.Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	return netClient
}
