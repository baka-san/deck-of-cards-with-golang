package main

import "fmt"

func main() {
	// Make a deck
	deck := newDeck()
	// deck.print()

	// Save the deck to a file
	// deck.saveToFile("my-deck")

	// Make another deck
	// deck2 := newDeckFromFile("my-deck")
	// deck2.print()

	// Deal a hand
	hand := deck.deal(5)
	hand.print()
	deck.print()

	// fmt.Println(deck.toString())

}

func printLine() {
	fmt.Println("-------------")
}
