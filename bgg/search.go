package bgg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SearchType string

type SearchFrame struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Total      string   `xml:"total,attr"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Search     []Search `xml:"item"`
}

type Search struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
	ID   string `xml:"id,attr"`
	Name struct {
		Text  string `xml:",chardata"`
		Type  string `xml:"type,attr"`
		Value string `xml:"value,attr"`
	} `xml:"name"`
	Yearpublished struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"yearpublished"`
}

const (
	Boardgame          SearchType = "boardgame"
	BoardgameAccessory SearchType = "boardgameaccessory"
	BoardgameExpansion SearchType = "boardgameexpansion"
	Videogame          SearchType = "videogame"
	RpgItem            SearchType = "rpgitem"
)

// Search returns a search result
func (client *Client) Search(query string, types []SearchType) ([]Search, error) {
	query = strings.Replace(query, " ", "+", -1)

	url := fmt.Sprintf("%s/searchFrame?query=%s", baseUrl, query)

	if types != nil && len(types) > 0 {
		url += "&type="
		for _, t := range types {
			url += string(t) + ","
		}
		// trim trailing comma
		url = url[:len(url)-1]
	}

	// Make the request to bgg
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", response.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Unmarshal to struct
	searchFrame := &SearchFrame{}
	err = xml.Unmarshal(body, searchFrame)
	if err != nil {
		return nil, err
	}

	return searchFrame.Search, nil
}
