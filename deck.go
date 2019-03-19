package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type Deck []string

// Make a Deck
func newDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print all cards in Deck
func (deck Deck) print() {
	for index, card := range deck {
		fmt.Println(index, card)
	}

	fmt.Println("-------------")
}

// Deal a hand
func (deck *Deck) deal(handSize int) Deck {
	// Shuffle deck
	deck.shuffle()

	// Select X cards from deck for hand
	hand := (*deck)[:handSize]

	// Remove said cards from deck
	*deck = (*deck)[handSize:]

	return hand
}

// Save the deck to a file
func (deck Deck) saveToFile(filename string) error {
	// Convert type
	data := []byte(deck.toString())

	return ioutil.WriteFile(filename, data, 0666)
}

func (deck Deck) toString() string {
	return strings.Join([]string(deck), ", ")
}

// Read in a deck from a file
func newDeckFromFile(filename string) Deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice), ", ")
	return Deck(stringSlice)
}

func (deck Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	n := rand.New(source)

	for i := range deck {
		j := n.Intn(len(deck) - 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}
