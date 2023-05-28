package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create new type Deck whic is slice of strings
type deck []string

func newDeck() deck {
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Q", "K"}

	cards := deck{}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func newDeckFromFile(filename string) deck {
	byteSlice, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toCsvString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toCsvString()), 0666)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[newPos], d[i] = d[i], d[newPos]
	}
}
