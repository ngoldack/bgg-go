package bgg

type Client struct{}

const (
	baseUrl string = "https://boardgamegeek.com/xmlapi2"
)

// New creates a new bgg client
func New() *Client {
	return &Client{}
}
