package forecast

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/rbo13/go-openweather/client"
	"github.com/rbo13/go-openweather/coords"
)

// ForecastData represents the
// forecast data from openweather
type ForecastData struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  float64 `json:"pressure"`
			SeaLevel  float64 `json:"sea_level"`
			GrndLevel float64 `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
		Rain  struct {
			ThreeH float64 `json:"3h"`
		} `json:"rain,omitempty"`
		Snow struct {
			ThreeH float64 `json:"3h"`
		} `json:"snow,omitempty"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country string `json:"country"`
	} `json:"city"`
}

type DailyForecastData struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	City    struct {
		GeonameID  int     `json:"geoname_id"`
		Name       string  `json:"name"`
		Lat        float64 `json:"lat"`
		Lon        float64 `json:"lon"`
		Country    string  `json:"country"`
		Iso2       string  `json:"iso2"`
		Type       string  `json:"type"`
		Population int     `json:"population"`
	} `json:"city"`
	Cnt  int `json:"cnt"`
	List []struct {
		Dt   int `json:"dt"`
		Temp struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity int     `json:"humidity"`
		Weather  []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Speed  float64 `json:"speed"`
		Deg    int     `json:"deg"`
		Clouds int     `json:"clouds"`
		Snow   float64 `json:"snow,omitempty"`
	} `json:"list"`
}

// Forecast ...
type Forecast struct {
	client.Client
	URL          string
	ForecastData *ForecastData
}

// NewForecast ...
func NewForecast(baseURL string) *Forecast {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &Forecast{
		Client: client.Client{
			APIKey: os.Getenv("OPENWEATHER_API_KEY"),
			HTTPClient: &http.Client{
				Timeout:   10 * time.Second,
				Transport: netTransport,
			},
		},
		URL:          baseURL,
		ForecastData: new(ForecastData),
	}
}

func (f *Forecast) GetByCityName(cityName string) (*ForecastData, error) {
	apiURL := fmt.Sprintf(f.URL+"/forecast?q=%s&appid=%s", cityName, f.Client.APIKey)
	err := f.Client.Request("GET", apiURL, &f.ForecastData)

	if err != nil {
		return nil, err
	}
	return f.ForecastData, nil
}

func (f *Forecast) GetByCityID(cityID int64) (*ForecastData, error) {
	apiURL := fmt.Sprintf(f.URL+"/forecast?id=%d&appid=%s", cityID, f.Client.APIKey)
	err := f.Client.Request("GET", apiURL, &f.ForecastData)

	if err != nil {
		return nil, err
	}
	return f.ForecastData, nil
}

func (f *Forecast) GetByCoordinates(coords coords.Coordinates) (*ForecastData, error) {
	apiURL := fmt.Sprintf(f.URL+"/forecast?lat=%g&lon=%g&appid=%s", coords.Latitude, coords.Longitude, f.Client.APIKey)
	err := f.Client.Request("GET", apiURL, &f.ForecastData)

	if err != nil {
		return nil, err
	}
	return f.ForecastData, nil
}

func (f *Forecast) GetByZipCode(zipCode, countryCode string) (*ForecastData, error) {
	apiURL := fmt.Sprintf(f.URL+"/forecast?zip=%s,%s&appid=%s", zipCode, countryCode, f.Client.APIKey)
	err := f.Client.Request("GET", apiURL, &f.ForecastData)

	if err != nil {
		return nil, err
	}
	return f.ForecastData, nil
}
