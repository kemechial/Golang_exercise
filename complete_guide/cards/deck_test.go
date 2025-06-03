package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got '%v'", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', but got '%v'", d[len(d)-1])	
	}


}	

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	filename := "_decktesting"
	defer os.Remove(filename) // Clean up after test

	d := newDeck()
	if err := d.saveToFile(filename); err != nil {
		t.Errorf("Failed to save deck to file: %v", err)
	}

	loadedDeck, err := newDeckFromFile(filename)
	if err != nil {
		t.Errorf("Failed to load deck from file: %v", err)
	}

	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck length of 52, but got %v", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got '%v'", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', but got '%v'", loadedDeck[len(loadedDeck)-1])	
	}
}