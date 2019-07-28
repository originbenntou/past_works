package main

import (
	"github.com/google/go-querystring/query"
	backlog "github.com/griffin-stewie/go-backlog"
	"net/url"
)

type customClient struct {
	backlog.Client
}

// NewClient returns Backlog HTTP Client
func newClient(baseURL *url.URL, APIKey string) (c customClient) {
	c.BaseURL = baseURL
	c.APIKey = APIKey

	return
}

func (c *customClient) Values() (url.Values, error) {
	return query.Values(c)
}
