package bgg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct{}

type SearchType string

const (
	baseUrl string = "https://boardgamegeek.com/xmlapi2"

	Boardgame          SearchType = "boardgame"
	BoardgameAccessory SearchType = "boardgameaccessory"
	BoardgameExpansion SearchType = "boardgameexpansion"
	Videogame          SearchType = "videogame"
	RpgItem            SearchType = "rpgitem"
)

// New creates a new bgg client
func New() *Client {
	return &Client{}
}

// Search returns a search result
func (cl *Client) Search(query string, types []SearchType, exact *int) (*Search, error) {
	query = strings.Replace(query, " ", "+", -1)

	url := fmt.Sprintf("%s/search?query=%s", baseUrl, query)

	if types != nil && len(types) > 0 {
		url += "&type="
		for _, t := range types {
			url += string(t) + ","
		}
		// trim trailing comma
		url = url[:len(url)-2]
	}

	if exact != nil {
		url += fmt.Sprintf("&exact=%v", *exact)
	}

	// Make the request to bgg
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Unmarshal to struct
	search := &Search{}
	err = xml.Unmarshal(body, search)
	if err != nil {
		return nil, err
	}

	return search, nil
}
