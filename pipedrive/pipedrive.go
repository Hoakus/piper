package pipedrive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	// Default consts
	defaultBaseURL = "pipedrive.com/api/"

	hostProtocol = "https"

	apiVersion = "2"

	// The amount of requests current API token can perform for the 10 seconds window.
	headerRateLimit = "X-RateLimit-Limit"

	// The amount of requests left for the 10 seconds window.
	headerRateRemaining = "X-RateLimit-Remaining"

	// The amount of seconds before the limit resets.
	headerRateReset = "X-RateLimit-Reset"
)

type Config struct {
	DomainName string
	APIKey     string
}

type Client struct {
	client  *http.Client
	BaseURL *url.URL
	apiKey  string
}

func (c *Client) createRequestURL(endpoint string, options *any) (string, error) {
	// "https://domain.pipedrive.com/api/v2/"
	reqURL, err := c.BaseURL.Parse(hostProtocol + "//:" + defaultBaseURL + "v" + apiVersion + "/")

	if err != nil {
		return endpoint, err
	}

	reqURL.Path += endpoint

	// If no options provided, assign the API key and return
	if options == nil {
		paramaters := url.Values{}
		paramaters.Add("api_key", c.apiKey)

		reqURL.RawQuery = paramaters.Encode()

		return reqURL.String(), nil
	}

	userQueries, err := query.Values(options)

	if err != nil {
		return endpoint, err
	}

	userQueries.Add("api_key", c.apiKey)
	reqURL.RawQuery = userQueries.Encode()

	return reqURL.String(), nil
}

func (c *Client) newRequest(method, endpoint string, options *any, body *any) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("Client.BaseURL must end with '/' : %v", c.BaseURL)
	}

	url, err := c.createRequestURL(endpoint, options)

	if err != nil {
		return nil, err
	}

	var buf io.Reader

	// Prepare payload if body exists
	if body != nil {
		buf := bytes.NewBuffer([]byte{})
		encoder := json.NewEncoder(buf)
		encoder.SetEscapeHTML(false)

		err = encoder.Encode(body)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)

	if err != nil {
		return nil, err
	}

	// Set the headers if we're sending a payload
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func NewClient(cfg *Config) *Client {
	baseURL, _ := url.Parse(cfg.DomainName + "." + defaultBaseURL)

	c := &Client{
		client:  http.DefaultClient,
		BaseURL: baseURL,
		apiKey:  cfg.APIKey,
	}
	return c
}
