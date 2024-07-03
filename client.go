package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/husseinamine/whatsapp/types"
)

type Client struct {
	WaHTTP     *http.Client
	BusinessID string
	Token      string
}

func (c *Client) SendMessage(message *types.Message) (interface{}, error) {
	uri, err := url.ParseRequestURI(
		fmt.Sprintf("%s/%s/messages", APIUrl, c.BusinessID),
	)

	if err != nil {
		return nil, err
	}

	// Write headers
	header := http.Header{}
	header.Add("Content-Type", ContentType)
	header.Add("Authorization", "Bearer "+c.Token)

	// Write body
	body, _ := json.Marshal(message)
	bodyReader := bytes.NewReader(body)
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

func NewClient(businessID string, token string) *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	return &Client{
		WaHTTP:     httpClient,
		BusinessID: businessID,
		Token:      token,
	}
}
