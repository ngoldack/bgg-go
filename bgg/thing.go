package bgg

import "encoding/xml"

type Thing struct {
	XMLName    xml.Name `xml:"items" json:"items,omitempty"`
	Text       string   `xml:",chardata" json:"text,omitempty"`
	Termsofuse string   `xml:"termsofuse,attr" json:"termsofuse,omitempty"`
	Item       struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Type string `xml:"type,attr" json:"type,omitempty"`
		ID   string `xml:"id,attr" json:"id,omitempty"`
		Name struct {
			Text      string `xml:",chardata" json:"text,omitempty"`
			Type      string `xml:"type,attr" json:"type,omitempty"`
			Sortindex string `xml:"sortindex,attr" json:"sortindex,omitempty"`
			Value     string `xml:"value,attr" json:"value,omitempty"`
		} `xml:"name" json:"name,omitempty"`
		Description   string `xml:"description"`
		Yearpublished struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Value string `xml:"value,attr" json:"value,omitempty"`
		} `xml:"yearpublished" json:"yearpublished,omitempty"`
		Minplayers struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Value string `xml:"value,attr" json:"value,omitempty"`
		} `xml:"minplayers" json:"minplayers,omitempty"`
		Maxplayers struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Value string `xml:"value,attr" json:"value,omitempty"`
		} `xml:"maxplayers" json:"maxplayers,omitempty"`
		Poll []struct {
			Text       string `xml:",chardata" json:"text,omitempty"`
			Name       string `xml:"name,attr" json:"name,omitempty"`
			Title      string `xml:"title,attr" json:"title,omitempty"`
			Totalvotes string `xml:"totalvotes,attr" json:"totalvotes,omitempty"`
			Results    []struct {
				Text       string `xml:",chardata" json:"text,omitempty"`
				Numplayers string `xml:"numplayers,attr" json:"numplayers,omitempty"`
				Result     []struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					Value    string `xml:"value,attr" json:"value,omitempty"`
					Numvotes string `xml:"numvotes,attr" json:"numvotes,omitempty"`
					Level    string `xml:"level,attr" json:"level,omitempty"`
				} `xml:"result" json:"result,omitempty"`
			} `xml:"results" json:"results,omitempty"`
		} `xml:"poll" json:"poll,omitempty"`
		Playingtime struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Value string `xml:"value,attr" json:"value,omitempty"`
		} `xml:"playingtime" json:"playingtime,omitempty"`
		Minplaytime struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Value string `xml:"value,attr" json:"value,omitempty"`
		} `xml:"minplaytime" json:"minplaytime,omitempty"`
		Maxplaytime struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Value string `xml:"value,attr" json:"value,omitempty"`
		} `xml:"maxplaytime" json:"maxplaytime,omitempty"`
		Minage struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Value string `xml:"value,attr" json:"value,omitempty"`
		} `xml:"minage" json:"minage,omitempty"`
		Link []struct {
			Text    string `xml:",chardata" json:"text,omitempty"`
			Type    string `xml:"type,attr" json:"type,omitempty"`
			ID      string `xml:"id,attr" json:"id,omitempty"`
			Value   string `xml:"value,attr" json:"value,omitempty"`
			Inbound string `xml:"inbound,attr" json:"inbound,omitempty"`
		} `xml:"link" json:"link,omitempty"`
	} `xml:"item" json:"item,omitempty"`
}
