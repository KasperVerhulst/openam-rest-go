package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Client struct {

	//HTTP client used to call openam Rest API
	HttpClient *http.Client

	//Admin token
	Token string

	//hostname of OpenAM server
	Host *url.URL

	//user-agent name
	UserAgent string
}

//create new go client
func NewClient(token string, host *url.URL) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		Token:      token,
		Host:       host,
	}
}

//create a new HTTP request to the OpenAM server
func (c *Client) createRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {

	var req *http.Request

	//create full url + path
	u, err := c.Host.Parse(path)
	if err != nil {
		return nil, err
	}

	switch method {
	case http.MethodGet, http.MethodDelete:

		req, err = http.NewRequest(method, u.String(), nil)

		if err != nil {
			return nil, err
		}

	case http.MethodPost, http.MethodPut:
		//check body is JSON
		var buf bytes.Buffer
		if body != nil {
			err := json.NewEncoder(&buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

		req, err = http.NewRequest(method, u.String(), &buf)

		if err != nil {
			return nil, err
		}

	default:
		log.Println("HTTP method is not supporterd")
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/json")
	return req, nil

}

//Execute the Rest API call
func (c *Client) doRequest()
