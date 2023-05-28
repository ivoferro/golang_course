package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	length := len(d)
	expectedLength := 52
	if length != expectedLength {
		t.Errorf("expected deck len of %v, but got %v", expectedLength, length)
	}

	firstCard := d[0]
	expectedFirstCard := "Ace of Spades"
	if firstCard != expectedFirstCard {
		t.Errorf("expected first card to be %v, but got %v", expectedFirstCard, firstCard)
	}

	lastCard := d[len(d)-1]
	expectedLastCard := "K of Clubs"
	if lastCard != expectedLastCard {
		t.Errorf("expected last card to be %v, but got %v", expectedLastCard, lastCard)
	}
}

func TestSaveAndLoadDeck(t *testing.T) {
	filename := "_deckTestFile"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)
	length := len(loadedDeck)
	expectedLength := 52
	if length != expectedLength {
		t.Errorf("expected deck len of %v, but got %v", expectedLength, length)
	}

	os.Remove(filename)
}
