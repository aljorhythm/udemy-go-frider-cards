package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d, e := newDeck(), 16
	if len(d) != e {
		t.Errorf("Expected deck length of %v, got %v", e, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", d[0])
	}
}

func TestSaveDeckAndReadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.writeToFile("_decktesting")
	loadedDeck, err := readDeckFromFile("_decktesting")

	if err != nil {
		t.Errorf("Failed reading")
	}
	if len(loadedDeck) != len(d) {
		t.Errorf("Loaded deck does not have same length")
	}
}
