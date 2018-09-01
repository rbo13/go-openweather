package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Client represents
// the client request to
// openweather api
type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

func (c *Client) Request(method, url string, data interface{}) error {

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

	return c.HTTPClient.Do(request)
}
