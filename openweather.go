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
// Sample: https://samples.openweathermap.org/data/2.5/weather?q=London,uk&appid=b6907d289e10d714a6e88b30761fae22
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
// Sample: https://samples.openweathermap.org/data/2.5/weather?id=2172797&appid=b6907d289e10d714a6e88b30761fae22
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
// Sample: https://samples.openweathermap.org/data/2.5/weather?lat=35&lon=139&appid=b6907d289e10d714a6e88b30761fae22
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
// weather by a given zip code and country code.
// If `countryCode` is not specified, it defaults to 'US',
// see: https://openweathermap.org/current#zip
// Sample: https://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b6907d289e10d714a6e88b30761fae22
func (c *Client) GetWeatherByZipCode(zipCode, countryCode string) (*WeatherData, error) {
	var weatherData WeatherData
	apiURL := fmt.Sprintf(baseURL+"/weather?zip=%s,%s&appid=%s", zipCode, countryCode, c.apiKey)

	err := c.request("GET", apiURL, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

// GetForecastByCityName returns the
// forecast by a given city
// Sample: https://samples.openweathermap.org/data/2.5/forecast?q=London,us&appid=b6907d289e10d714a6e88b30761fae22
func (c *Client) GetForecastByCityName(cityName string) (*ForecastData, error) {
	var forecastData ForecastData
	apiURL := fmt.Sprintf(baseURL+"/forecast?q=%s&appid=%s", cityName, c.apiKey)

	err := c.request("GET", apiURL, &forecastData)

	if err != nil {
		return nil, err
	}

	return &forecastData, nil
}

// GetForecastByCityID returns the
// forecast by a given city id
// Sample: https://samples.openweathermap.org/data/2.5/forecast?id=524901&appid=b6907d289e10d714a6e88b30761fae22
func (c *Client) GetForecastByCityID(cityID string) (*ForecastData, error) {
	var forecastData ForecastData
	apiURL := fmt.Sprintf(baseURL+"/forecast?id=%s&appid=%s", cityID, c.apiKey)

	err := c.request("GET", apiURL, &forecastData)

	if err != nil {
		return nil, err
	}

	return &forecastData, nil
}

// GetForecastByCoordinates returns the
// forecast by a given coordinates
// Sample: https://samples.openweathermap.org/data/2.5/forecast?lat=35&lon=139&appid=b6907d289e10d714a6e88b30761fae22
func (c *Client) GetForecastByCoordinates(coords Coordinates) (*ForecastData, error) {
	var forecastData ForecastData
	apiURL := fmt.Sprintf(baseURL+"/forecast?lat=%g&lon=%g&appid=%s", coords.Latitude, coords.Longitude, c.apiKey)

	err := c.request("GET", apiURL, &forecastData)

	if err != nil {
		return nil, err
	}
	return &forecastData, nil
}

// GetForecastByZipCode returns the
// forecast by a given zip code and country code.
// If country code is not specied, it defaults to 'USA',
// for reference, see: https://openweathermap.org/forecast5#zip
// Sample: https://samples.openweathermap.org/data/2.5/forecast?zip=94040&appid=b6907d289e10d714a6e88b30761fae22

func (c *Client) GetForecastByZipCode(zipCode, countryCode string) (*ForecastData, error) {
	var forecastData ForecastData
	apiURL := fmt.Sprintf(baseURL+"/forecast?zip=%s,%s&appid=%s", zipCode, countryCode, c.apiKey)

	err := c.request("GET", apiURL, &forecastData)

	if err != nil {
		return nil, err
	}

	return &forecastData, nil
}

// GetDailyForecastByCityName returns the
// daily forecast by a given city name, unit
// and count/frequency of the forecast.
// Sets default value for unit and count if
// none is specified.
func (c *Client) GetDailyForecastByCityName(cityName, unit, count string) (*DailyForecastData, error) {
	var dailyForecastData DailyForecastData

	if unit == "" {
		unit = "metric"
	}

	if count == "" {
		count = "7" // defaults to 7 days of forecast
	}

	apiURL := fmt.Sprintf(baseURL+"/forecast/daily?q=%s&units=%s&cnt=%s&appid=%s", cityName, unit, count, c.apiKey)

	err := c.request("GET", apiURL, &dailyForecastData)

	if err != nil {
		return nil, err
	}

	return &dailyForecastData, nil
}

// GetDailyForecastByCityID returns the
// daily forecast by a given city id and count/frequency
// of the forecast. Sets default value to count if
// none is specified.
func (c *Client) GetDailyForecastByCityID(cityID, count string) (*DailyForecastData, error) {
	var dailyForecastData DailyForecastData

	if count == "" {
		count = "7" // defaults to 7 days of forecast
	}

	apiURL := fmt.Sprintf(baseURL+"/forecast/daily?id=%s&cnt=%s&appid=%s", cityID, count, c.apiKey)

	err := c.request("GET", apiURL, &dailyForecastData)

	if err != nil {
		return nil, err
	}

	return &dailyForecastData, nil
}

// GetDailyForecastByCoordinates returns the
// daily forecast by a given city id and count/frequency
// of the forecast. Sets default value to count if
// none is specified.
func (c *Client) GetDailyForecastByCoordinates(coords Coordinates, count string) (*DailyForecastData, error) {
	var dailyForecastData DailyForecastData

	if count == "" {
		count = "16" // defaults to 16 days of forecast
	}

	apiURL := fmt.Sprintf(baseURL+"/forecast/daily?lat=%g&lon=%g&cnt=%s&appid=%s", coords.Latitude, coords.Longitude, count, c.apiKey)

	err := c.request("GET", apiURL, &dailyForecastData)

	if err != nil {
		return nil, err
	}

	return &dailyForecastData, nil
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
