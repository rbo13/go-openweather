package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Coordinates defines the cordinates of a place
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Client represents
// the client request to
// openweather api
type Client struct {
	apiKey string
}

const baseURL string = "https://api.openweathermap.org/data/2.5"

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
	var weatherData WeatherData
	apiURL := fmt.Sprintf(baseURL+"/weather?q=%s&appid=%s", cityName, c.apiKey)
	err := request("GET", apiURL, &weatherData)
	if err != nil {
		return nil, err
	}
	return &weatherData, nil
}

// GetWeatherByCityID returns the
// weather in a given cityID
func (c *Client) GetWeatherByCityID(cityID int64) (*WeatherData, error) {
	var weatherData WeatherData
	apiURL := fmt.Sprintf(baseURL+"/weather?id=%d&appid=%s", cityID, c.apiKey)
	err := request("GET", apiURL, &weatherData)
	if err != nil {
		return nil, err
	}
	return &weatherData, nil
}

// GetWeatherByCoordinates returns the
// weather by a given coordinates
func (c *Client) GetWeatherByCoordinates(coords Coordinates) (*WeatherData, error) {
	var weatherData WeatherData
	apiURL := fmt.Sprintf(baseURL+"/weather?lat=%g&lon=%g&appid=%s", coords.Latitude, coords.Longitude, c.apiKey)

	err := request("GET", apiURL, &weatherData)
	if err != nil {
		return nil, err
	}
	return &weatherData, nil
}

func request(method, url string, data interface{}) error {
	client := buildHTTPClient()

	resp, err := buildHTTPRequest(method, url, client)

	if err != nil && resp.Body == nil {
		return err
	}

	defer resp.Body.Close()
	buffer, err := ioutil.ReadAll(resp.Body)

	if err != nil && buffer == nil {
		return err
	}

	return json.Unmarshal(buffer, &data)
}

func buildHTTPClient() *http.Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	return client
}

func buildHTTPRequest(method, url string, client *http.Client) (*http.Response, error) {
	request, err := http.NewRequest(method, url, nil)

	if err != nil && request.Body == nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")

	return client.Do(request)
}
