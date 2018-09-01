package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

// Client represents
// the client request to
// openweather api
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// Weather ...
type Weather struct {
	Client
	Url         string
	WeatherData *WeatherData
}

// Forecast ...
type Forecast struct {
	Client
	Url          string
	ForecastData *ForecastData
}

const baseURL string = "https://api.openweathermap.org/data/2.5"

func NewWeather(baseUrl string) *Weather {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &Weather{
		Client: Client{
			apiKey: os.Getenv("OPENWEATHER_API_KEY"),
			httpClient: &http.Client{
				Timeout:   10 * time.Second,
				Transport: netTransport,
			},
		},
		Url:         baseUrl,
		WeatherData: new(WeatherData),
	}
}

func NewForecast(baseUrl string) *Forecast {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &Forecast{
		Client: Client{
			apiKey: os.Getenv("OPENWEATHER_API_KEY"),
			httpClient: &http.Client{
				Timeout:   10 * time.Second,
				Transport: netTransport,
			},
		},
		Url:          baseUrl,
		ForecastData: new(ForecastData),
	}
}

func (w *Weather) GetByCityName(cityName string) (*WeatherData, error) {
	apiUrl := fmt.Sprintf(w.Url+"/weather?q=%s&appid=%s", cityName, w.Client.apiKey)
	err := w.Client.request("GET", apiUrl, &w.WeatherData)

	if err != nil {
		return nil, err
	}
	return w.WeatherData, nil
}

func (w *Weather) GetByCityID(cityID int64) (*WeatherData, error) {
	apiUrl := fmt.Sprintf(w.Url+"/weather?id=%d&appid=%s", cityID, w.Client.apiKey)
	err := w.Client.request("GET", apiUrl, &w.WeatherData)

	if err != nil {
		return nil, err
	}
	return w.WeatherData, nil
}

func (w *Weather) GetByCoordinates(coords Coordinates) (*WeatherData, error) {
	apiUrl := fmt.Sprintf(w.Url+"/weather?lat=%g&lon=%g&appid=%s", coords.Latitude, coords.Longitude, w.Client.apiKey)
	err := w.Client.request("GET", apiUrl, &w.WeatherData)

	if err != nil {
		return nil, err
	}

	return w.WeatherData, nil
}

func (w *Weather) GetByZipCode(zipCode, countryCode string) (*WeatherData, error) {
	apiUrl := fmt.Sprintf(w.Url+"/weather?zip=%s,%s&appid=%s", zipCode, countryCode, w.Client.apiKey)
	err := w.Client.request("GET", apiUrl, &w.WeatherData)

	if err != nil {
		return nil, err
	}
	return w.WeatherData, nil
}

func (f *Forecast) GetByCityName(cityName string) (*ForecastData, error) {
	apiUrl := fmt.Sprintf(f.Url+"/forecast?q=%s&appid=%s", cityName, f.Client.apiKey)
	err := f.Client.request("GET", apiUrl, &f.ForecastData)

	if err != nil {
		return nil, err
	}
	return f.ForecastData, nil
}

func (f *Forecast) GetByCityID(cityID int64) (*ForecastData, error) {
	apiUrl := fmt.Sprintf(f.Url+"/forecast?id=%d&appid=%s", cityID, f.Client.apiKey)
	err := f.Client.request("GET", apiUrl, &f.ForecastData)

	if err != nil {
		return nil, err
	}
	return f.ForecastData, nil
}

func (f *Forecast) GetByCoordinates(coords Coordinates) (*ForecastData, error) {
	apiUrl := fmt.Sprintf(f.Url+"/forecast?lat=%g&lon=%g&appid=%s", coords.Latitude, coords.Longitude, f.Client.apiKey)
	err := f.Client.request("GET", apiUrl, &f.ForecastData)

	if err != nil {
		return nil, err
	}
	return f.ForecastData, nil
}

func (f *Forecast) GetByZipCode(zipCode, countryCode string) (*ForecastData, error) {
	apiUrl := fmt.Sprintf(f.Url+"/forecast?zip=%s,%s&appid=%s", zipCode, countryCode, f.Client.apiKey)
	err := f.Client.request("GET", apiUrl, &f.ForecastData)

	if err != nil {
		return nil, err
	}
	return f.ForecastData, nil
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
