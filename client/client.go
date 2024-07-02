package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	WaHTTP       *http.Client
	DefaultPhone string
	Token        string
}

func (c *Client) SendRequest(body string) (interface{}, error) {
	uri, err := url.ParseRequestURI(
		fmt.Sprintf("%s/%s/messages", APIUrl, c.DefaultPhone),
	)

	if err != nil {
		return nil, err
	}

	// Write headers
	header := http.Header{}
	header.Add("Content-Type", ContentType)
	header.Add("Authorization", "Bearer "+c.Token)

	// Write body
	bodyReader := strings.NewReader(body)
	bodyCloser := io.NopCloser(bodyReader)

	// Do request
	res, err := c.WaHTTP.Do(&http.Request{
		Method: "POST",
		URL:    uri,
		Header: header,
		Body:   bodyCloser,
	})

	if err != nil {
		return nil, err
	}

	// Read response body
	resBody := make(map[string]interface{})
	resBodyDecoder := json.NewDecoder(res.Body)

	resBodyDecoder.Decode(&resBody)

	return resBody, nil
}

func NewClient(defaultPhone string, token string) *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	return &Client{
		WaHTTP:       httpClient,
		DefaultPhone: defaultPhone,
		Token:        token,
	}
}
