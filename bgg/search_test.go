package bgg_test

import (
	"github.com/ngoldack/bgg-go/bgg"
	"testing"
)

func TestSearch(t *testing.T) {
	client := bgg.New()
	t.Run("test with query only", func(t *testing.T) {
		searches, err := client.Search("Chess", nil)
		if err != nil {
			t.Error(err)
		}
		if len(searches) == 0 {
			t.Errorf("expected multipe results, got %d", len(searches))
		}
	})

	t.Run("test for single type", func(t *testing.T) {
		searchType := bgg.Boardgame

		searchResults, err := client.Search("Chess", []bgg.SearchType{searchType})
		if err != nil {
			t.Error(err)
		}

		if len(searchResults) == 0 {
			t.Errorf("expected multipe results, got %d", len(searchResults))
		}

		for _, search := range searchResults {
			if search.Type != string(searchType) {
				t.Errorf("expected search result %s to have type %s, got %s", search.ID, searchType, search.Type)
			}
		}
	})

	t.Run("test for multiple types", func(t *testing.T) {
		types := []bgg.SearchType{
			bgg.Boardgame,
			bgg.BoardgameExpansion,
		}

		searchResults, err := client.Search("Carcassonne", types)
		if err != nil {
			t.Error(err)
		}

		if len(searchResults) == 0 {
			t.Errorf("expected multipe results, got %d", len(searchResults))
		}

		for _, search := range searchResults {
			switch search.Type {
			case string(bgg.Boardgame):
			case string(bgg.BoardgameExpansion):
			default:
				t.Errorf("expected search result type to be %s or %s, got %s", bgg.Boardgame, bgg.BoardgameExpansion, search.Type)
			}
		}
	})
}
