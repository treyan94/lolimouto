package openweathermap

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const baseUrl = "https://api.openweathermap.org/data/2.5/weather"

type Client struct {
	apiKey     string
	httpClient http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		httpClient: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (cl Client) Search(q string) (c *City) {
	req, err := searchReq(q, cl)
	if err != nil {
		return
	}

	res, getErr := cl.httpClient.Do(req)
	if getErr != nil {
		return
	}

	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return
	}

	_ = json.Unmarshal(body, &c)

	return
}

func searchReq(q string, cl Client) (*http.Request, error) {
	queryParams := url.Values{
		"q":     {q},
		"appid": {cl.apiKey},
	}

	return http.NewRequest(
		http.MethodGet,
		baseUrl+"?"+queryParams.Encode(),
		nil,
	)
}
