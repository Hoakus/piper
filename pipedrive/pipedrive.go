package pipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	hostProtocol   = "https"
	defaultBaseURL = "pipedrive.com/"

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
	apiKey  string
	BaseURL *url.URL
	Rate    *Rate

	// Every pipedrive module is a piper
	common piper

	// TODO: Add additional modules
	Organization *OrganizationPiper
	Task         *TaskPiper
	Activities   *ActivitiesPiper
	Leads        *LeadsPiper
	DealFields   *DealFieldsPiper
	Persons      *PersonsPiper
}

type piper struct {
	client *Client
}

type Rate struct {
	Limit     int
	Remaining int
	Reset     time.Duration
	mu        sync.Mutex
}

func (c *Client) setRateDetails(response *http.Response) error {
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

func (c *Client) createRequestURL(endpoint string, params any) (string, error) {
	reqURL := *c.BaseURL

	reqURL.Path += endpoint

	if params == nil {
		return reqURL.String(), nil
	}

	userQueries, err := query.Values(params)

	if err != nil {
		return endpoint, err
	}

	reqURL.RawQuery = userQueries.Encode()

	return reqURL.String(), nil
}

func checkForErrors(response *http.Response) error {
	if response.StatusCode < 299 {
		return nil
	}

	ed := ErrorDetails{}

	json.NewDecoder(response.Body).Decode(&ed)

	errorResponse := ErrorResponse{
		Response: response,
		Details:  ed,
	}

	return errorResponse
}

func (c *Client) NewRequest(method, endpoint string, params, body any) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("Client.BaseURL must end with '/' : %v", c.BaseURL)
	}

	url, err := c.createRequestURL(endpoint, params)

	if err != nil {
		return nil, err
	}

	var payload io.ReadWriter
	if body != nil {
		buf := &bytes.Buffer{}
		encoder := json.NewEncoder(buf)
		encoder.SetEscapeHTML(false)

		err = encoder.Encode(body)

		if err != nil {
			return nil, err
		}
		payload = buf
	}

	request, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new http.Request: %v", err)
	}

	request.Header.Set("x-api-token", c.apiKey)

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return request, nil
}

func (c *Client) Do(ctx context.Context, request *http.Request, container any) (*http.Response, error) {
	if c.Rate != nil {
		c.Rate.mu.Lock()

		if c.Rate.Remaining == 0 {
			// Add's 1 back to the current client Rate
			// to prevent blocking on next request once error
			// has been handled by the caller
			c.Rate.Remaining = 1

			c.Rate.mu.Unlock()

			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:

				return nil, errors.New("Rate limit reached. Aborted")
			}

		} else {
			c.Rate.mu.Unlock()
		}
	}

	request = request.WithContext(ctx)

	response, err := c.client.Do(request)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, fmt.Errorf("Failed to do http.Request: %v", err)
		}
	}

	err = c.setRateDetails(response)

	if err != nil {
		log.Fatalf("FATAL: Could not set rate details for pipedrive client: %v", err)
	}

	err = checkForErrors(response)

	if err != nil {
		return nil, fmt.Errorf("Client received non-200 status code: %v", err)
	}

	// Reading a small, fixed amount (like 512 bytes) ensures that any
	// initial buffered data is consumed, helping the HTTP transport to
	// cleanly reset the connection for reuse.
	defer func() {
		io.CopyN(io.Discard, response.Body, 512)
		response.Body.Close()
	}()

	// Unmarshal the data to the provided response struct
	err = json.NewDecoder(response.Body).Decode(container)

	if err != nil {
		return nil, fmt.Errorf("Failed to decode response into container: %v", err)
	}

	return response, nil
}

func NewClient(cfg *Config) *Client {
	urlString := fmt.Sprintf("%s://%s.%s", hostProtocol, cfg.DomainName, defaultBaseURL)
	baseURL, _ := url.Parse(urlString)

	// Initialize rate with 1 request to ensure getRateDetails updates the rate
	// from the first response before allowing additional requests
	rate := &Rate{
		Limit:     0,
		Remaining: 1,
		Reset:     time.Duration(0),
	}

	c := &Client{
		client:  http.DefaultClient,
		BaseURL: baseURL,
		apiKey:  cfg.APIKey,
		Rate:    rate,
	}

	c.common.client = c

	// this allows functionality for every piper to use the
	// underlying logic of client
	c.Organization = (*OrganizationPiper)(&c.common)
	c.Task = (*TaskPiper)(&c.common)
	c.Activities = (*ActivitiesPiper)(&c.common)
	c.Leads = (*LeadsPiper)(&c.common)
	c.DealFields = (*DealFieldsPiper)(&c.common)
	c.Persons = (*PersonsPiper)(&c.common)

	return c
}
