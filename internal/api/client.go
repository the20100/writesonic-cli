package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseURL = "https://api.writesonic.com/v2/business/content"

// Client is the Writesonic API client.
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new authenticated Writesonic API client.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// Post sends an authenticated POST request to the given path with query params
// and a JSON body. Returns the raw response bytes.
func (c *Client) Post(path string, queryParams url.Values, body map[string]interface{}) ([]byte, error) {
	endpoint := baseURL + path + "?" + queryParams.Encode()

	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, reqBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("X-API-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode == 422 {
		var ve ValidationError
		if json.Unmarshal(data, &ve) == nil && len(ve.Detail) > 0 {
			return nil, &ve
		}
		return nil, fmt.Errorf("validation error (422): %s", string(data))
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(data))
	}

	return data, nil
}

// PostResults sends a POST and decodes the response into a slice of ContentResult.
func (c *Client) PostResults(path string, queryParams url.Values, body map[string]interface{}) ([]ContentResult, error) {
	data, err := c.Post(path, queryParams, body)
	if err != nil {
		return nil, err
	}
	var results []ContentResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return results, nil
}

// PostLandingPages sends a POST and decodes into a slice of LandingPage.
func (c *Client) PostLandingPages(queryParams url.Values, body map[string]interface{}) ([]LandingPage, error) {
	data, err := c.Post("/landing-pages", queryParams, body)
	if err != nil {
		return nil, err
	}
	var results []LandingPage
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return results, nil
}
