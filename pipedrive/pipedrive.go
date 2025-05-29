package pipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultBaseURL = "pipedrive.com/"

	hostProtocol = "https"

	// The amount of requests current API token can perform for the 10 seconds window.
	headerRateLimit = "X-RateLimit-Limit"

	// The amount of requests left for the 10 seconds window.
	headerRateRemaining = "X-RateLimit-Remaining"

	// The amount of seconds before the limit resets.
	headerRateReset = "X-RateLimit-Reset"
)

// Config must be passed in to call of pipedrive.NewClient()
type Config struct {
	DomainName string
	APIKey     string
}

// Client that is returned from call of pipedrive.NewClient()
type Client struct {
	client  *http.Client
	BaseURL *url.URL
	apiKey  string
	Rate    *RateDetails

	// Every pipedrive module is a piper
	common piper

	// TODO: Add additional modules
	Organization *OrganizationPiper
}

type piper struct {
	client *Client
}

type RateDetails struct {
	Limit     int
	Remaining int
	Reset     time.Duration
	mu        sync.Mutex
}

type Response struct {
	Status       HTTPStatus
	HttpResponse *http.Response
}

type HTTPStatus struct {
	Code int
	Text string
}

func (c *Client) setRateDetails(response *http.Response) error {
	// sets the clients Rate field to the values received
	// from the most recent response in Client.Do -
	// probably need a more elegant solutions for this.
	rateLimit, err := strconv.Atoi(response.Header.Get(headerRateLimit))
	if err != nil {
		return err
	}

	rateRemaining, err := strconv.Atoi(response.Header.Get(headerRateRemaining))
	if err != nil {
		return err
	}

	rateReset, err := time.ParseDuration(response.Header.Get(headerRateReset) + "s")
	if err != nil {
		return err
	}

	c.Rate.mu.Lock()
	defer c.Rate.mu.Unlock()

	c.Rate.Limit = rateLimit
	c.Rate.Remaining = rateRemaining
	c.Rate.Reset = rateReset

	return nil
}

func (c *Client) createRequestURL(endpoint string, queryParams any, apiVersion string) (string, error) {
	reqURL := *c.BaseURL

	// some pipedrive endpoints require endpoint requests to contain
	// information in the URL, but not as a query param. In these cases,
	// the correct value needs to be passed in as part of the endpoint
	// string, not queryParams
	reqURL.Path += "api" + "/" + "v" + apiVersion + "/" + endpoint

	if queryParams == nil {
		return reqURL.String(), nil
	}

	// uses go-querystrings to add paramaters to the final query
	userQueries, err := query.Values(queryParams)

	if err != nil {
		return endpoint, err
	}

	reqURL.RawQuery = userQueries.Encode()

	return strings.ToLower(reqURL.String()), nil
}

func (c *Client) NewRequest(method, endpoint, apiVersion string, queryParams, body any) (*http.Request, error) {
	// apiVersion must be passed in as not all endpoints have migrated to v2
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("Client.BaseURL must end with '/' : %v", c.BaseURL)
	}

	url, err := c.createRequestURL(endpoint, queryParams, apiVersion)

	if err != nil {
		return nil, err
	}

	// Prepare payload if body exists
	var buf io.Reader

	if body != nil {
		buf := bytes.NewBuffer([]byte{})
		encoder := json.NewEncoder(buf)
		encoder.SetEscapeHTML(false)

		err = encoder.Encode(body)

		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, url, buf)

	if err != nil {
		return nil, err
	}

	// Set the API key in the headers
	// pipedrive no longer accepts the apiKey as a query param
	request.Header.Set("x-api-token", c.apiKey)

	// Set the headers if we're sending a payload
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return request, nil
}

func (c *Client) Do(ctx context.Context, request *http.Request, container any) (*Response, error) {
	// Check Rate Limiting first
	if c.Rate != nil {
		c.Rate.mu.Lock()

		if c.Rate.Remaining == 0 {
			c.Rate.mu.Unlock()

			select {
			// wait for the rate limit to reset before continuing
			case <-time.After(c.Rate.Reset * time.Second):
			// unless the context is done first
			case <-ctx.Done():
				{
					return nil, ctx.Err()
				}
			}

		} else {
			c.Rate.mu.Unlock()
		}
	}

	// starting request with context here
	request = request.WithContext(ctx)

	response, err := c.client.Do(request)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}

	var httpStatus HTTPStatus
	httpStatus.Code = response.StatusCode
	httpStatus.Text = response.Status

	// Separates status from the response for easier use
	// for the end user. Still returns the full httpResponse
	// for debugging errors
	newResponse := &Response{Status: httpStatus, HttpResponse: response}

	// set the Rate details returned from the response
	// so we can control request flow on subsequent calls and prevent
	// errors from rate limitng
	err = c.setRateDetails(response)
	if err != nil {
		fmt.Println(err)
	}

	if newResponse.Status.Code > 299 {
		return newResponse, fmt.Errorf("Request failed: %v - URL: %v", newResponse.Status.Text, request.URL)
	}

	// Reading a small, fixed amount (like 512 bytes) ensures that any
	// initial buffered data is consumed, helping the HTTP transport to
	// cleanly reset the connection for reuse.
	defer func() {
		io.CopyN(io.Discard, response.Body, 512)
		response.Body.Close()
	}()

	// Unmarshal the data to the provided piperResponse struct
	// passed in from the piper.method call
	err = json.NewDecoder(response.Body).Decode(container)

	if err != nil {
		return newResponse, err
	}

	return newResponse, nil
}

func NewClient(cfg *Config) *Client {
	baseURLString := hostProtocol + "://" + cfg.DomainName + "." + defaultBaseURL
	baseURL, _ := url.Parse(baseURLString)

	rate := &RateDetails{
		Limit:     0,
		Remaining: 1,
		Reset:     time.Duration(0),
	}

	newClient := &Client{
		client:  http.DefaultClient,
		BaseURL: baseURL,
		apiKey:  cfg.APIKey,
		Rate:    rate,
	}

	// newClient.Organization.client = newClient
	newClient.common.client = newClient

	// this allows functionality for every piper to use the
	// underlying logic of client
	newClient.Organization = (*OrganizationPiper)(&newClient.common)

	return newClient
}
