package get_issue_comment

import (
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
