package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) DoRequest(method, path, payload, params string, headers map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s%s", c.BaseURL, path, params)
	var req *http.Request
	var err error

	if method == "POST" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	// Add headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	log.Printf("Sending %s request to %s", method, url)
	client := &http.Client{}
	return client.Do(req)
}
