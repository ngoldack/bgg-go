package bgg

import "net/http"

type Client struct {
	httpClient *http.Client
}

const (
	baseUrl string = "https://boardgamegeek.com/xmlapi2"
)

// New creates a new bgg client
func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}
