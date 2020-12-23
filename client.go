package bitfinex

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client holds details that allow communication with the Bitfinex API.
type Client struct {
	Host       string
	HTTPClient *http.Client
}

// New returns a new Client.
func New() *Client {
	t := time.Second * 2

	return &Client{
		Host:       "https://api-pub.bitfinex.com",
		HTTPClient: &http.Client{Timeout: t},
	}
}

// Tickers returns the details for the given ticker symbols.
//
// See https://docs.bitfinex.com/reference#rest-public-tickers
func (c *Client) Tickers(pairs []string) ([]Ticker, error) {
	path := fmt.Sprintf("/v2/tickers?symbols=%s", strings.Join(pairs, ","))

	body, err := c.get(path)
	if err != nil {
		return nil, err
	}

	tickers, err := ParseTickers(body)
	if err != nil {
		return nil, err
	}

	return tickers, nil
}

// get a response from a URL.
//
// This method will handle closing off the body.
func (c Client) get(path string) ([]byte, error) {
	url := c.Host + path

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
