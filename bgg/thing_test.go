package bgg_test

import (
	"github.com/ngoldack/bgg-go/bgg"
	"testing"
)

var testGameIds = []string{
	"309116",
	"173346",
	"316380",
}

func TestThings(t *testing.T) {
	client := bgg.New()

	t.Run("test no id", func(t *testing.T) {
		things, err := client.Thing(nil, "")
		if err != nil {
			t.Error(err)
		}
		if len(things) != 0 {
			t.Errorf("expected no things, got %v", len(things))
		}
	})

	t.Run("test single id", func(t *testing.T) {
		things, err := client.Thing(nil, testGameIds[:1]...)
		if err != nil {
			t.Error(err)
		}
		if len(things) != 1 {
			t.Errorf("expected 1 thing, got %v", len(things))
		}
	})

	t.Run("test multiple ids", func(t *testing.T) {
		things, err := client.Thing(nil, testGameIds...)
		if err != nil {
			t.Error(err)
		}
		if len(things) != 3 {
			t.Errorf("expected 3 things, got %v", len(things))
		}
	})
}
