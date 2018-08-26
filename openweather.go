package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// Client represents
// the client request to
// openweather api
type Client struct {
	apiKey     string
	httpClient *http.Client
}

const baseURL string = "https://api.openweathermap.org/data/2.5"

// https://samples.openweathermap.org/data/2.5/weather?q=London,uk&appid=b6907d289e10d714a6e88b30761fae22

// NewClient returns the Client struct
func NewClient(apiKey string) *Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout:   10 * time.Second,
			Transport: netTransport,
		},
	}
}

// GetWeatherByCityName returns the
// weather in a given city
func (c *Client) GetWeatherByCityName(cityName string) (*WeatherData, error) {
	var weatherData WeatherData
	apiURL := fmt.Sprintf(baseURL+"/weather?q=%s&appid=%s", cityName, c.apiKey)
	err := c.request("GET", apiURL, &weatherData)
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
	err := c.request("GET", apiURL, &weatherData)
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

	err := c.request("GET", apiURL, &weatherData)
	if err != nil {
		return nil, err
	}
	return &weatherData, nil
}

// GetWeatherByZipCode returns the
// weather by a given zip code and country code
func (c *Client) GetWeatherByZipCode(zipCode, countryCode string) (*WeatherData, error) {
	var weatherData WeatherData
	apiURL := fmt.Sprintf(baseURL+"/weather?zip=%s,%s&appid=%s", zipCode, countryCode, c.apiKey)

	err := c.request("GET", apiURL, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func (c *Client) request(method, url string, data interface{}) error {

	resp, err := c.buildHTTPRequest(method, url)

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

func (c *Client) buildHTTPRequest(method, url string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, nil)

	if err != nil && request.Body == nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")

	return c.httpClient.Do(request)
}
